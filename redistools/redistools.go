package redistools

import (
	"errors"
	"github.com/go-redis/redis"
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
	_, err := rc.Client.Ping().Result()
	if nil != err {
		return err
	}
	return nil
}

// Lock 加锁
// outTime 加锁时长(秒)
// rslt 成功为 true
func (rc *RedisClient) Lock(key string, outTime int) (ok bool, err error) {
	return rc.Client.SetNX(key, "lock", time.Duration(outTime)*time.Second).Result()
}

//Unlock 解锁
func (rc *RedisClient) UnLock(key string) (err error) {
	return rc.Client.Del(key).Err()
}
