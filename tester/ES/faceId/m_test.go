package main

import(
	"encoding/json"
	"fmt"
	"testing"
	

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/faceId"
)
/**
*获取实人认证连接,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/ucnkx4
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	getVerifyTokenRequest := faceId.GetVerifyTokenRequest{}
    getVerifyTokenRequest.BizNo = "117577928278938"
	getVerifyTokenRequest.ReturnUrl = "https://www.es.com/api/realNameAuthReturn"
	getVerifyTokenRequest.NotifyUrl = "https://www.es.com/api/realNameAuthNotice"
	getVerifyTokenRequest.PageBgColor = "#eeeeee"
	getVerifyTokenRequest.OcrOnly = "true"
	getVerifyTokenRequest.RetIdImg = "true"
	getVerifyTokenRequest.OcrIncIdBack = "true"
	getVerifyTokenRequest.ReturnImg = "false"
	getVerifyTokenRequest.ReturnLiveVideo = "true"
	bizContent := getVerifyTokenRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(faceId.GetVerifyTokenRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	getVerifyTokenResponse := faceId.GetVerifyTokenResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &getVerifyTokenResponse)
	fmt.Println("服务器返回业务数据 : ", getVerifyTokenResponse.ToJsonStr())
}