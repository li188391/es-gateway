package main

import(
	"encoding/json"
	"fmt"
	"testing"
	
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/es2"
	
	
)
/**
*一键登录,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/ggqpgh
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	getMobileRequest := es2.GetMobileRequest{}
    getMobileRequest.AppName = "测试app"
	getMobileRequest.DeviceModel = "HUAWEI-Z132"
	getMobileRequest.PackageId = "com.es.test"
	getMobileRequest.Platform = "1"
	getMobileRequest.Token = "701478a1e3b4b7e3978ea69469410f13"

	bizContent := getMobileRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(es2.GetMobileRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	getMobileResponse := es2.GetMobileResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &getMobileResponse)
	fmt.Println("服务器返回业务数据 : ", getMobileResponse.ToJsonStr())
}