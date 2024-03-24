package main

import(
	"encoding/json"
	"fmt"
	"testing"
	
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/zszn"
	
	
)
/**
*运营商三要素校验,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/phim68
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	idNamePhoneCheckRequest := zszn.IdNamePhoneCheckRequest{}
    idNamePhoneCheckRequest.IdName = "李翠花"
	idNamePhoneCheckRequest.IdNumber = "596223555698546855"
	idNamePhoneCheckRequest.Mobile = "16599998888"
	

	bizContent := idNamePhoneCheckRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(zszn.IdNamePhoneCheckRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	idNamePhoneCheckResponse := zszn.IdNamePhoneCheckResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &idNamePhoneCheckResponse)
	fmt.Println("服务器返回业务数据 : ", idNamePhoneCheckResponse.ToJsonStr())
}