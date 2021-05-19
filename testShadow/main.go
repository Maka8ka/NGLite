package main

import (
	"NknGontrol/module/nkntransfer"
)

func main() {
	// nkntransfer.StartListener(config.Controllerid)
	nkntransfer.RsaCompressListener("testshadowtrap")

}

// // Dealmsg 消息处理函数
// type Dealmsg func(interface{})

// // Dealmsg 该函数类型实现了一个方法
// func Dealmsg(msg *nkn.Message) {
// 	fmt.Println("Receive message from", msg.Src+":", string(msg.Data))
// 	msg.Reply([]byte("recevie ok ,hello"))
// }
