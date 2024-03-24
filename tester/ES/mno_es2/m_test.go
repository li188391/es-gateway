package main

import(
	"encoding/json"
	"fmt"
	"testing"
	
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/mno_es2"
	
	
)
/**
*本机号码认证,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/ggpzrn
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	verifyMobileRequest := mno_es2.VerifyMobileRequest{}
    verifyMobileRequest.AppName = "测试app"
	verifyMobileRequest.DeviceModel = "HUAWEI-Z132"
	verifyMobileRequest.PackageId = "com.es.test"
	verifyMobileRequest.PhoneNum = "13588995478"
	verifyMobileRequest.Platform = "1"
	verifyMobileRequest.Token = "701478a1e3b4b7e3978ea69469410f13"

	bizContent := verifyMobileRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(mno_es2.VerifyMobileRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	verifyMobileResponse := mno_es2.VerifyMobileResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &verifyMobileResponse)
	fmt.Println("服务器返回业务数据 : ", verifyMobileResponse.ToJsonStr())
}