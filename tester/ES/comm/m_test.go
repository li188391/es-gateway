package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/comm"
	"com.es.gateway.sdk/sdk/util"
)

func Test1(t *testing.T) {
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	esBalanceRequest := comm.EsBalanceRequest{}
	esBalanceRequest.TransId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	bizContent := esBalanceRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(comm.EsBalanceRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	esBalanceResponse := comm.EsBalanceResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &esBalanceResponse)
	fmt.Println("服务器返回业务数据 : ", esBalanceResponse.ToJsonStr())
}
