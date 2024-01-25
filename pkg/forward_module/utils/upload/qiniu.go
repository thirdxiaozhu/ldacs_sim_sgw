package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/forward_module/f_global"
)

type Qiniu struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Qiniu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: f_global.GVA_CONFIG.Qiniu.Bucket}
	mac := qbox.NewMac(f_global.GVA_CONFIG.Qiniu.AccessKey, f_global.GVA_CONFIG.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		global.LOGGER.Error("function file.Open() failed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close()                                                  // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		global.LOGGER.Error("function formUploader.Put() failed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() failed, err:" + putErr.Error())
	}
	return f_global.GVA_CONFIG.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(f_global.GVA_CONFIG.Qiniu.AccessKey, f_global.GVA_CONFIG.Qiniu.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(f_global.GVA_CONFIG.Qiniu.Bucket, key); err != nil {
		global.LOGGER.Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: qiniuConfig
//@description: 根据配置文件进行返回七牛云的配置
//@return: *storage.Config

func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      f_global.GVA_CONFIG.Qiniu.UseHTTPS,
		UseCdnDomains: f_global.GVA_CONFIG.Qiniu.UseCdnDomains,
	}
	switch f_global.GVA_CONFIG.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
