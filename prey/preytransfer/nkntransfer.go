package nkntransfer

import (
	"encoding/hex"
	"fmt"

	config "NknGontrol/module"

	"github.com/nknorg/nkn-sdk-go"
)

var Account *nkn.Account

func init() {
	seed, err := hex.DecodeString(config.Secretseed)
	if err != nil {
		fmt.Println(err)
	}
	Account, _ = nkn.NewAccount(seed)
	if err != nil {
		fmt.Println(err)
	}
		fmt.Println(Account.Seed(), Account.PubKey())
}

func StartListener(listenid string) {

	listener, err := nkn.NewMultiClient(Account, listenid, config.TransThreads, false, nil)
	if err != nil {
		fmt.Println(err)
	}
		fmt.Println(listener.Address())
	<-listener.OnConnect.C
		fmt.Println("Connection opened.")

		for {
		msg := <-listener.OnMessage.C

		fmt.Println("Receive message from", msg.Src+":", string(msg.Data))
		msg.Reply([]byte("recevie ok ,hello")) 	}

}

func Sender(sourceid string, destinationid string, content interface{}) (interface{}, error) {
	source, err := nkn.NewMultiClient(Account, sourceid, config.TransThreads, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	destination, err := nkn.NewMultiClient(Account, destinationid, config.TransThreads, false, nil)
	if err != nil {
		fmt.Println(err)
	}
		fmt.Println(source.Address())
	<-source.OnConnect.C
		fmt.Println("Connection opened.")

	response, err := source.Send(nkn.NewStringArray(destination.Address()), content, nil)
					return response.Next().Data, err
}
