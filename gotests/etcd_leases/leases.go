package main

import (
	"context"
	"fmt"
	"log"

	"github.com/coreos/etcd/clientv3"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2380"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// minimum lease TTL is 5-second
	lease1, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	lease2, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatal(err)
	}

	// after 5 seconds, the key 'foo' will be removed
	_, err = cli.Put(context.TODO(), "foo", "bar", clientv3.WithLease(lease1.ID))
	fmt.Println("put key 'foo' with lease", lease1.ID)
	fmt.Println("lease1: ", lease1)
	if err != nil {
		log.Fatal(err)
	}

	response, err := cli.Get(context.TODO(), "foo")
	fmt.Println("response: ", response)
	if err != nil {
		log.Fatal(err)
	}
	key := response.Kvs[0]

	fmt.Printf("got key %s with lease %d\n", key.Key, key.Lease)

	_, err = cli.Put(context.TODO(), "foo", "bar", clientv3.WithLease(lease2.ID))
	fmt.Println("put key 'foo' with lease", lease2.ID)
	if err != nil {
		log.Fatal(err)
	}

	response, err = cli.Get(context.TODO(), "foo")
	if err != nil {
		log.Fatal(err)
	}
	key = response.Kvs[0]
	fmt.Printf("got key %s with lease %d\n", key.Key, key.Lease)
}
