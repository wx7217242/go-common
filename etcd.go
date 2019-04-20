package util

import (
	etcd_client "github.com/coreos/etcd/clientv3"
	"time"
)

type EtcdConf struct {
	Addr    string
	Timeout int
}

func InitEtcd(conf EtcdConf) (client *etcd_client.Client, err error) {
	client, err = etcd_client.New(etcd_client.Config{
		Endpoints:   []string{conf.Addr},
		DialTimeout: time.Duration(conf.Timeout) * time.Second,
	})

	return
}
