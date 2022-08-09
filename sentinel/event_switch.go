package sentinel

import (
	"context"
	"fmt"
	"log"
	"sentinel/modify"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var ipMaster = ""
var i = 0
var dataFile modify.DataFile
var arrDataFile []modify.DataFile

func newsub() *redis.SentinelClient {
	sentinel := redis.NewSentinelClient(&redis.Options{
		Addr: ":26376",
	})
	return sentinel
}

func newChange(arrCmd []string, ip string) {
	modify.ChangeConfig(modify.GetIPConfig(arrCmd), "REDIS_HOST: '"+ip+"'")
	ipMaster = ip
	for i := 0; i < len(arrCmd); i++ {
		dataFile.IP = ipMaster
		dataFile.PATH = arrCmd[i]
		arrDataFile = append(arrDataFile, dataFile)
	}
}

func oldChange(ip string) {
	modify.ChangeConfig(arrDataFile, "REDIS_HOST: '"+ip+"'")
	ipMaster = ip
	for i := 0; i < len(arrDataFile); i++ {
		arrDataFile[i].IP = ipMaster
	}
}

func NewMaster(arrCmd []string) {
	sub := newsub()
	//pub := sentinel.Subscribe(ctx, "+switch-master")
	for {
		pub := sub.Subscribe(ctx, "+switch-master")
		msg, err := pub.ReceiveMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}
		ch := pub.Channel()
		fmt.Println("Channel", ch)
		fmt.Println("Msg Channel" + msg.Channel)
		fmt.Println("Msg String" + msg.String())
		fmt.Println("Msg Pattern " + msg.Pattern)
		fmt.Println("Msg Payload " + msg.Payload)
		addr, err := sub.GetMasterAddrByName(ctx, "mymaster").Result()
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			newChange(arrCmd, addr[0])
			i++
		} else {
			oldChange(addr[0])
			//i++
		}
	}
}
