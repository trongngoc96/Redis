package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewMaster() {
	sentinel := redis.NewSentinelClient(&redis.Options{
		Addr: ":26376",
	})
	pub := sentinel.Subscribe(ctx, "+switch-master")
	msg, err := pub.ReceiveMessage(ctx)
	ch := pub.Channel()
	fmt.Println("Channel", ch)
	fmt.Println("Channel" + msg.Channel)
	fmt.Println("Msg String" + msg.String())
	fmt.Println("Msg Pattern " + msg.Pattern)
	fmt.Println("Msg Payload " + msg.Payload)
	addr, err := sentinel.GetMasterAddrByName(ctx, "mymaster").Result()
	ChangeConfig(addr)
	if err != nil {
	}
	fmt.Println("Ip new master", addr)
	NewMaster()
}

func main() {
	NewMaster()
	fmt.Scanln()
}
