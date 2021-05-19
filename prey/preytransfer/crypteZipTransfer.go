package nkntransfer

import (
	config "NknGontrol/module"
	"NknGontrol/module/cipher"
	"NknGontrol/module/compress"
	"fmt"

	"github.com/nknorg/nkn-sdk-go"
)

func RsaCompressSender(sourceid string, destinationid string, content interface{}) (interface{}, error) {
	contentcrypt, err := cipher.RsaEncrypt(content.([]byte), []byte(config.RsaPublicKey))
	if err != nil {
		fmt.Println(err)
	}
	contentzip, err := compress.Zip(contentcrypt)
	if err != nil {
		fmt.Println(err)
	}
	response, _ := Sender(sourceid, destinationid, contentzip.Bytes())
	return response, err
}

func RsaCompressListener(listenid string) {

	listener, err := nkn.NewMultiClient(Account, listenid, config.TransThreads, false, nil)
	if err != nil {
		fmt.Println(err, "listen data error")
	}
		fmt.Println(listener.Address())
	<-listener.OnConnect.C
		fmt.Println("Connection opened.")

		for {
		msg := <-listener.OnMessage.C

		unzipdata, err := compress.UnZip(msg.Data) 		if err != nil {
			fmt.Println(err, "unzip data error")
		}
		realdata, err := cipher.RsaDecrypt(unzipdata, []byte(config.RsaPrivateKey)) 		if err != nil {
			fmt.Println(err, "unzip data error")
		}
		fmt.Println(realdata)

		fmt.Println("Receive message from", msg.Src+":", string(realdata))
		msg.Reply([]byte("recevie ok ,hello")) 	}

}

func AesCompressSender(sourceid string, destinationid string, content interface{}) (interface{}, error) {
	contentcrypt, err := cipher.AesCbcEncrypt(content.([]byte), []byte(config.AesKey))
	if err != nil {
		fmt.Println(err)
	}
	contentzip, err := compress.Zip(contentcrypt)
	if err != nil {
		fmt.Println(err)
	}
	response, _ := Sender(sourceid, destinationid, contentzip.Bytes())
	return response, err
}

func AesCompressListener(b []byte) []byte {
	unzip, err := compress.UnZip(b)
	if err != nil {
		fmt.Println("接受数据无法解码", err)
	}
	data, _ := cipher.AesCbcDecrypt(unzip, []byte(config.AesKey))
	return data
}

