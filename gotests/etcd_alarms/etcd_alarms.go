package main

import (
	"context"
	"fmt"
	"log"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/embed"
	"github.com/coreos/etcd/etcdserver/etcdserverpb"
)

func main() {
	cfg := embed.NewConfig()
	cfg.Dir = "."
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer etcd.Close()

	cli, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2380"},
	})
	defer cli.Close()

	resp, _ := cli.MemberList(context.Background())

	member := resp.Members[0]

	alarmRequest := &etcdserverpb.AlarmRequest{
		Action:   etcdserverpb.AlarmRequest_ACTIVATE,
		MemberID: member.ID,
		Alarm:    3,
	}

	etcd.Server.Alarm(context.Background(), alarmRequest)
	alarms := etcd.Server.Alarms()
	fmt.Println("fetching alarms")
	fmt.Println(alarms)
}
