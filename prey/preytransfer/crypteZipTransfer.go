package nkntransfer

import (
	config "NknGontrol/module"
	"NknGontrol/module/cipher"
	"NknGontrol/module/compress"
	"fmt"

	"github.com/nknorg/nkn-sdk-go"
)

// RsaCompressSender 压缩并加密发送数据
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

// RsaCompressListener 解压解密读取数据
func RsaCompressListener(listenid string) {

	listener, err := nkn.NewMultiClient(Account, listenid, config.TransThreads, false, nil)
	if err != nil {
		fmt.Println(err, "listen data error")
	}
	//debug 1 line
	fmt.Println(listener.Address())
	<-listener.OnConnect.C
	//debug 1 line
	fmt.Println("Connection opened.")

	// 监听循环取出数据处理
	for {
		msg := <-listener.OnMessage.C

		unzipdata, err := compress.UnZip(msg.Data) //解压数据
		if err != nil {
			fmt.Println(err, "unzip data error")
		}
		realdata, err := cipher.RsaDecrypt(unzipdata, []byte(config.RsaPrivateKey)) //解密数据
		if err != nil {
			fmt.Println(err, "unzip data error")
		}
		fmt.Println(realdata)

		fmt.Println("Receive message from", msg.Src+":", string(realdata))
		msg.Reply([]byte("recevie ok ,hello")) //传入interface类型，可以是byte数据，也可以是string
	}

}

// AesCompressSender 压缩并加密发送数据
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

//AesCompressListener 解压缩并解密
func AesCompressListener(b []byte) []byte {
	unzip, err := compress.UnZip(b)
	if err != nil {
		fmt.Println("接受数据无法解码", err)
	}
	data, _ := cipher.AesCbcDecrypt(unzip, []byte(config.AesKey))
	return data
}

