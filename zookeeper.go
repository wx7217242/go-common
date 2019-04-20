package util

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type ZooKeeperConf struct {
	servers []string
	timeout int
}

type ZooKeeperClient struct {
	Conn   *zk.Conn
	Events <-chan zk.Event
}

func InitZooKeeper(conf ZooKeeperConf) (cli *ZooKeeperClient, err error) {

	conn, events, err := zk.Connect(conf.servers, time.Duration(conf.timeout)*time.Second)
	if err != nil {
		return nil, err
	}

	cli = &ZooKeeperClient{
		Conn:   conn,
		Events: events,
	}

	return
}
