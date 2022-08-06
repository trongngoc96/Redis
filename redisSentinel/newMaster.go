package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ChangeConfig(addr string) {
	data, err := ioutil.ReadFile("config.js")
	if err != nil {
		fmt.Println(err)
	}
	//REDIS_HOST: '10.22.7.107',
	fmt.Print(string(data))

	f, err := os.OpenFile("config.js", os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	re := regexp.MustCompile(`REDIS_HOST..............'`)
	if _, err = f.WriteString(

		re.ReplaceAllString(string(data), "REDIS_HOST:"+" '"+addr+"' ")); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("config.js")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}

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
	justString := strings.Join(addr, " ")
	ChangeConfig(justString)
	if err != nil {
	}
	fmt.Println("Ip new master", addr)
	NewMaster()
}

func main() {
	NewMaster()
	fmt.Scanln()
}
