package redistools

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	Client *redis.Client
	Conf   *RedisConf
}

type RedisConf struct {
	Addr     string
	Password string
	DB       int
}

//链接Redis
func (rc *RedisClient) ConnRedis() error {
	if len(rc.Conf.Addr) == 0 {
		return errors.New("redis address is null")
	}
	//connect redis
	rc.Client = redis.NewClient(&redis.Options{
		Addr:     rc.Conf.Addr,
		Password: rc.Conf.Password,
		DB:       rc.Conf.DB,
	})
	//check redis connect ready
	var ctx = context.Background()
	_, err := rc.Client.Ping(ctx).Result()
	if nil != err {
		return err
	}
	return nil
}

// Lock 加锁
// outTime 加锁时长(秒)
// ok 成功为 true
func (rc *RedisClient) Lock(context context.Context, key string, outTime int) (ok bool, err error) {
	return rc.Client.SetNX(context, key, "lock", time.Duration(outTime)*time.Second).Result()
}

//Unlock 解锁
func (rc *RedisClient) UnLock(context context.Context, key string) (err error) {
	return rc.Client.Del(context, key).Err()
}
