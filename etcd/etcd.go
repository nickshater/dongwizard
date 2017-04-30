//Package etcd manages ETCD for the api settings
package etcd

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const (
	etcdStart string = "etcd"
	etcdStop  string = "pkill etcd"
)

//StartETCD starts the ETCD service
func StartETCD() {
	c := exec.Command(etcdStart)

	c.Start()
}

//KillETCD kills the ETCD service
func KillETCD() {
	c := exec.Command(etcdStop)

	c.Start()
}

//GetValue returns a string value matching the passed key
func GetValue(key string) string {
	StartETCD()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("etcd client error ", err)
	}
	defer cli.Close()

	requestTimeout := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Println("getvalue ctx error ", err)
	}
	var v string
	for _, ev := range resp.Kvs {
		v = string(ev.Value)
	}
	KillETCD()
	return v
}
