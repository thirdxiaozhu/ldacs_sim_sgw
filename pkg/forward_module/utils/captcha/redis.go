package captcha

import (
	"context"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/forward_module/f_global"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
		Context:    context.TODO(),
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	err := f_global.GVA_REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.LOGGER.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := f_global.GVA_REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.LOGGER.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := f_global.GVA_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.LOGGER.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
