package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/smsall"
	
)
/**
*发送国际短信，参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/xw4q62
**/
func Test1(t *testing.T) {
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	sendmsgAllRequest := smsall.SendmsgAllRequest{}
	sendmsgAllRequest.Mobile = "+8613666666666"
	sendmsgAllRequest.TemplateID = "20201024185922"
	sendmsgAllRequest.TemplateParamSet = "123, 889"
	sendmsgAllRequest.CallbackUrl = "http://test.dev.esandcloud.com"

	bizContent := sendmsgAllRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(smsall.SendmsgAllRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	sendmsgAllResponse := smsall.SendmsgAllResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &sendmsgAllResponse)
	fmt.Println("服务器返回业务数据 : ", sendmsgAllResponse.ToJsonStr())
}
