package common

import (
	etcd_client "go.etcd.io/etcd/clientv3"
	"time"
)

type EtcdConf struct {
	Endpoints []string
	Timeout   int
}

func InitEtcd(conf EtcdConf) (client *etcd_client.Client, err error) {
	client, err = etcd_client.New(etcd_client.Config{
		Endpoints:   conf.Endpoints,
		DialTimeout: time.Duration(conf.Timeout) * time.Second,
	})

	return
}
