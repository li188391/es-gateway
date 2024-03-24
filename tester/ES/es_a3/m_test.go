package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
	"encoding/base64"
	"io/ioutil"
	"log"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/es_a3"
	"com.es.gateway.sdk/sdk/util"
)
/**
*身份三要素校验，参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/ko3aof
**/
func Test1(t *testing.T) {
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	esA3AuthRequest := es_a3.EsA3AuthRequest{}
	esA3AuthRequest.TransId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	esA3AuthRequest.HadLiving = false
	esA3AuthRequest.IdNumber = "110101199003078312"
	esA3AuthRequest.Name = "张三"
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	esA3AuthRequest.PhotoData = base64Str

	bizContent := esA3AuthRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(es_a3.EsA3AuthRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	esA3AuthResponse := es_a3.EsA3AuthResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &esA3AuthResponse)
	fmt.Println("服务器返回业务数据 : ", esA3AuthResponse.ToJsonStr())
}