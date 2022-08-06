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
var ipMaster = ""
var i = 0

func getIPConfig() {
	data, err := ioutil.ReadFile("config.js")
	if err != nil {
		fmt.Println(err)
	}
	sRead := string(data)
	removeDistance := strings.ReplaceAll(sRead, " ", "")
	removeLine := strings.ReplaceAll(removeDistance, "\n", "")
	result2 := strings.Index(removeLine, "',REDIS_PORT")
	result1 := strings.Index(removeLine, `REDIS_HOST`)
	ipRedisOfConfig := removeLine[result1+12 : result2]
	ipMaster = ipRedisOfConfig
	fmt.Println("IP Of File Config.js", ipRedisOfConfig)
}

func ChangeConfig(dataFind string, addr string) {
	data, err := ioutil.ReadFile("config.js")
	if err != nil {
		fmt.Println(err)
	}
	//REDIS_HOST: '10.22.7.107',
	f, err := os.OpenFile("config.js", os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	re := regexp.MustCompile(`REDIS_HOST: ` + `'` + dataFind + `'`)
	if _, err = f.WriteString(
		re.ReplaceAllString(string(data), addr)); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("config.js")
	if err != nil {
		fmt.Println(err)
	}
}

func NewMaster() {
	sentinel := redis.NewSentinelClient(&redis.Options{
		Addr: ":26376",
	})
	pub := sentinel.Subscribe(ctx, "+switch-master")
	msg, err := pub.ReceiveMessage(ctx)
	ch := pub.Channel()
	fmt.Println("Channel", ch)
	fmt.Println("Msg Channel" + msg.Channel)
	fmt.Println("Msg String" + msg.String())
	fmt.Println("Msg Pattern " + msg.Pattern)
	fmt.Println("Msg Payload " + msg.Payload)
	addr, err := sentinel.GetMasterAddrByName(ctx, "mymaster").Result()
	if i == 0 {
		ChangeConfig(ipMaster, "REDIS_HOST: '"+addr[0]+"'")
		ipMaster = addr[0]
		i++
	} else {
		ChangeConfig(ipMaster, "REDIS_HOST: '"+addr[0]+"'")
		ipMaster = addr[0]
		if err != nil {
		}
		i++
	}
	NewMaster()
}

func main() {
	getIPConfig()
	NewMaster()
	fmt.Scanln()
}
