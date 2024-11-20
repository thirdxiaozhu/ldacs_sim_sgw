package backward_module

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hdt3213/godis/lib/logger"
	"github.com/hdt3213/godis/lib/sync/atomic"
	"github.com/hdt3213/godis/lib/sync/wait"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type ServiceHandler interface {
	Serve(msg []byte, id uint32)
	Close(id uint32)
}

// 客户端连接的抽象
type GscConn struct {
	// tcp 连接
	Id uint32

	Conn   net.Conn
	Server *SgwServer

	bufr *bufio.Reader
	buft *bufio.Writer
	// 当服务端开始发送数据时进入waiting, 阻止其它goroutine关闭连接
	// wait.Wait是作者编写的带有最大等待时间的封装:
	// https://github.com/HDT3213/godis/blob/master/src/lib/sync/wait/wait.go
	Waiting wait.Wait
}

type SgwServer struct {
	// 保存所有工作状态client的集合(把map当set用)
	// 需使用并发安全的容器
	activeConn sync.Map
	Addr       string
	Handler    ServiceHandler
	closing    atomic.Boolean // 关闭状态标识位
}

func (c *GscConn) Serve(msg []byte) {
	handler := c.Server.Handler
	handler.Serve(msg, c.Id)
}

func (c *GscConn) Close(id uint32) {
	handler := c.Server.Handler
	handler.Close(id)
}

func ListenAndServe(addr string, handler ServiceHandler) {
	server := &SgwServer{Addr: addr, Handler: handler}
	server.ListenAndServeWithSignal()
}

func (s *SgwServer) NewConn(conn net.Conn) *GscConn {
	c := &GscConn{
		Id:     uuid.New().ID(),
		Conn:   conn,
		Server: s,
	}
	return c
}

// ServeNewConnection 监听并提供服务，并在收到 closeChan 发来的关闭通知后关闭
func (s *SgwServer) ServeNewConnection(listener net.Listener, closeChan <-chan struct{}) {
	// 监听关闭通知
	go func() {
		<-closeChan
		logger.Info("shutting down...")
		// 停止监听，listener.Accept()会立即返回 io.EOF
		_ = listener.Close()
		// 关闭应用层服务器
		_ = s.Close()
	}()

	// 在异常退出后释放资源
	defer func() {
		// close during unexpected error
		_ = listener.Close()
		_ = s.Close()
	}()
	ctx := context.Background()
	var waitDone sync.WaitGroup
	for {
		// 监听端口, 阻塞直到收到新连接或者出现错误
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		c := s.NewConn(conn)
		// 开启 goroutine 来处理新连接
		logger.Info("accept link")
		waitDone.Add(1)
		go func() {
			defer func() {
				waitDone.Done()
			}()
			c.serve(ctx)
		}()
	}
	waitDone.Wait()
}

// ListenAndServeWithSignal 监听中断信号并通过 closeChan 通知服务器关闭
func (s *SgwServer) ListenAndServeWithSignal() {
	closeChan := make(chan struct{})
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigCh
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return
	}
	logger.Info(fmt.Sprintf("bind: %s, start listening...", s.Addr))
	s.ServeNewConnection(listener, closeChan)
}

func (c *GscConn) serve(ctx context.Context) {
	// 关闭中的 handler 不会处理新连接
	if c.Server.closing.Get() {
		_ = c.Conn.Close()
		return
	}

	//c.Server.activeConn.Store(c, struct{}{}) // 记住仍然存活的连接
	c.Server.activeConn.Store(c.Id, c) // 记住仍然存活的连接

	c.bufr = bufio.NewReader(c.Conn)
	c.buft = bufio.NewWriter(c.Conn)
	var msg [1024]byte
	for {
		n, err := c.bufr.Read(msg[:])
		if err != nil {
			if err == io.EOF {
				logger.Info("connection close")
				c.Close(c.Id)
				c.Server.activeConn.Delete(c.Id)
			} else {
				logger.Warn(err)
			}
			return
		}
		c.Serve(msg[:n])
	}
}

func (c *GscConn) SendPkt(pktJ []byte) {
	// 发送数据前先置为waiting状态，阻止连接被关闭
	c.Waiting.Add(1)

	c.Conn.Write(pktJ)
	c.Waiting.Done()
}

// 关闭客户端连接
func (c *GscConn) CloseConnection() error {
	// 等待数据发送完成或超时
	c.Waiting.WaitWithTimeout(10 * time.Second)
	_ = c.Conn.Close()
	return nil
}

// 关闭服务器
func (s *SgwServer) Close() error {
	logger.Info("handler shutting down...")
	s.closing.Set(true)
	// 逐个关闭连接
	s.activeConn.Range(func(key interface{}, val interface{}) bool {
		client := val.(*GscConn)
		_ = client.CloseConnection()
		return true
	})
	return nil
}
