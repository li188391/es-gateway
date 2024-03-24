package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/es_a2"
	"com.es.gateway.sdk/sdk/util"
)
/**
*身份二要素校验，参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/sus5d7
**/
func Test1(t *testing.T) {
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	esA2AuthRequest := es_a2.EsA2AuthRequest{}
	esA2AuthRequest.TransId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	esA2AuthRequest.IdNumber = "110101199003078312"
	esA2AuthRequest.Name = "张三"
	bizContent := esA2AuthRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(es_a2.EsA2AuthRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	esA2AuthResponse := es_a2.EsA2AuthResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &esA2AuthResponse)
	fmt.Println("服务器返回业务数据 : ", esA2AuthResponse.ToJsonStr())
}