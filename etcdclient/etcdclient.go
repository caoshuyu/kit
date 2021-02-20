package etcdclient

import (
	"context"
	"github.com/caoshuyu/kit/dlog"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var EtcdClient *clientv3.Client

var etcdConf = struct {
	endpoints     []string
	timeOutSecond int
	username      string
	password      string
}{}

func SetEndpoints(endpoints []string) {
	etcdConf.endpoints = endpoints
}

func AddEndpoint(endpoint string) {
	etcdConf.endpoints = append(etcdConf.endpoints, endpoint)
}

func SetUsername(username string) {
	etcdConf.username = username
}

func SetPassword(password string) {
	etcdConf.password = password
}

func getEndpoints() []string {
	if 0 == len(etcdConf.endpoints) {
		etcdConf.endpoints = append(etcdConf.endpoints, "127.0.0.1:2379")
	}
	return etcdConf.endpoints
}

func SetTimeOutSecond(timeOutSecond int) {
	etcdConf.timeOutSecond = timeOutSecond
}

func GetTimeOutSecond() int {
	if 0 == etcdConf.timeOutSecond {
		return 5
	}
	return etcdConf.timeOutSecond
}

func InitClient() {
	var err error
	cfg := clientv3.Config{
		Endpoints:   getEndpoints(),
		DialTimeout: time.Duration(GetTimeOutSecond()) * time.Second,
	}
	if len(etcdConf.username) > 0 {
		cfg.Username = etcdConf.username
	}
	if len(etcdConf.password) > 0 {
		cfg.Password = etcdConf.password
	}
	EtcdClient, err = clientv3.New(cfg)
	if nil != err {
		panic(err)
	}
}

func GetValue(key string) (value []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(GetTimeOutSecond())*time.Second)
	resp, err := EtcdClient.Get(ctx, key)
	cancel()
	if nil != err {
		dlog.ERROR("EtcdClient", "GetValue", "err", err)
		return
	}
	for _, ev := range resp.Kvs {
		value = ev.Value
		break
	}
	return
}

func WatchValue(key string, c chan<- []byte) {
	value := GetValue(key)
	c <- value
	watch := EtcdClient.Watch(context.Background(), key)
	for wresp := range watch {
		for _, v := range wresp.Events {
			c <- v.Kv.Value
		}
	}
}
