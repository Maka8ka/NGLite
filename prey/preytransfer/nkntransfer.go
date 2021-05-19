package nkntransfer

import (
	"encoding/hex"
	"fmt"

	config "NknGontrol/module"

	"github.com/nknorg/nkn-sdk-go"
)

// Account 全局变量
var Account *nkn.Account

// init 初始化Account
func init() {
	seed, err := hex.DecodeString(config.Secretseed)
	if err != nil {
		fmt.Println(err)
	}
	Account, _ = nkn.NewAccount(seed)
	if err != nil {
		fmt.Println(err)
	}
	//初始化打印种子和公钥
	fmt.Println(Account.Seed(), Account.PubKey())
}

// StartListener 函数传入需监听的id值
func StartListener(listenid string) {

	listener, err := nkn.NewMultiClient(Account, listenid, config.TransThreads, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	//debug 1 line
	fmt.Println(listener.Address())
	<-listener.OnConnect.C
	//debug 1 line
	fmt.Println("Connection opened.")

	// 监听循环取出数据处理
	for {
		msg := <-listener.OnMessage.C

		fmt.Println("Receive message from", msg.Src+":", string(msg.Data))
		msg.Reply([]byte("recevie ok ,hello")) //传入interface类型，可以是byte数据，也可以是string
	}

}

// Sender 函数 传入源地址id，目的地址id以及消息内容接口类型可以是byte[],也可以是string
func Sender(sourceid string, destinationid string, content interface{}) (interface{}, error) {
	source, err := nkn.NewMultiClient(Account, sourceid, config.TransThreads, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	destination, err := nkn.NewMultiClient(Account, destinationid, config.TransThreads, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	//debug 1 line
	fmt.Println(source.Address())
	<-source.OnConnect.C
	//debug 1 line
	fmt.Println("Connection opened.")

	response, err := source.Send(nkn.NewStringArray(destination.Address()), content, nil)
	// if err != nil {
	// 	fmt.Println(err, "sender response error")
	// }
	// fmt.Println(string(response.Next().Data), "response ok")
	return response.Next().Data, err
}
