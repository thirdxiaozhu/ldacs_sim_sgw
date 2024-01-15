package forward_module

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartGin() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	server := &http.Server{
		Addr:    ":7777",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}
