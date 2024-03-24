package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/sms"
	
)
/**
*发送短信，参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/fv3enr
**/
func Test1(t *testing.T) {
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	sendmsgRequest := sms.SendmsgRequest{}
	sendmsgRequest.Mobile = "13666666666"
	sendmsgRequest.TemplateID = "20201024185922"
	sendmsgRequest.TemplateParamSet = "123, 889"
	sendmsgRequest.CallbackUrl = "http://test.dev.esandcloud.com"

	bizContent := sendmsgRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(sms.SendmsgRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	sendmsgResponse := sms.SendmsgResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &sendmsgResponse)
	fmt.Println("服务器返回业务数据 : ", sendmsgResponse.ToJsonStr())
}