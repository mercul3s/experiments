package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {

	client, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	defer client.Close()
	lease, err := createLease(context.Background(), client, 15)
	if err != nil {
		panic(err)
	}

	res, _ := client.Put(context.TODO(), "/keyPrefix", "keyVal", clientv3.WithLease(lease))
	fmt.Println(res)
	response, _ := client.Get(context.TODO(), "/keyPrefix")

	kv := response.Kvs[0]
	fmt.Println(kv.Lease)
	leaseID := clientv3.LeaseID(kv.Lease)

	if err := getTTL(context.Background(), client, leaseID); err != nil {
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)
	err = extendLease(context.Background(), client, leaseID)
	if err != nil {
		panic(err)
	}

	if err := getTTL(context.Background(), client, leaseID); err != nil {
		fmt.Println(err)
	}
	time.Sleep(2 * time.Second)

	if err := getTTL(context.Background(), client, leaseID); err != nil {
		fmt.Println(err)
	}

	watchKey(context.Background(), client, "/keyPrefix")
	client.Close()
}

func createLease(ctx context.Context, cli *clientv3.Client, ttl int64) (clientv3.LeaseID, error) {
	lease, err := cli.Grant(context.Background(), 15)
	fmt.Printf("Initial Lease TTL: %v\n", lease.TTL)
	fmt.Printf("Lease ID: %v\n", lease.ID)
	fmt.Printf("Hex encoded lease ID: %x\n", lease.ID)
	if err != nil {
		return 0, err
	}
	return lease.ID, nil

}
func extendLease(ctx context.Context, cli *clientv3.Client, leaseID clientv3.LeaseID) error {
	ka, kaerr := cli.KeepAliveOnce(ctx, leaseID)
	if kaerr != nil {
		return kaerr
	}
	fmt.Println("extending lease ttl:", ka.TTL)
	return nil
}

func getTTL(ctx context.Context, cli *clientv3.Client, leaseID clientv3.LeaseID) error {
	ttl, err := cli.TimeToLive(context.Background(), leaseID)
	if err != nil {
		return err
	}
	fmt.Printf("Current lease TTL: %v\n", ttl.TTL)
	return nil
}

func watchKey(ctx context.Context, cli *clientv3.Client, key string) {
	rch := cli.Watch(ctx, key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type.String() == "DELETE" {
				fmt.Println("event deleted")
				return
			}
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
