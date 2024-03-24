package main

import(
	"encoding/json"
	"fmt"
	"testing"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/efaceid"
)
/**
*获取离线人脸认证SDK授权证书,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/vqhu4m
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	getCertificateRequest := efaceid.GetCertificateRequest{}
	getCertificateRequest.AppName = "一砂测试DEMO"
	getCertificateRequest.Platform = "ANDROID"
	getCertificateRequest.PackageID = "com.esandinfo.efaceid.demo"
	getCertificateRequest.DeviceID = "adfloiLJSDFADSDSFAIIWDFsSseeelphyhaoga"
	getCertificateRequest.SdkName = "haixin"
	getCertificateRequest.SdkVersion = "1.0"

	bizContent := getCertificateRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(efaceid.GetCertificateRequestAct,bizContent)
	if gatewayResponse.Code !="0000"{
		t.Error("服务器返回异常 : ",gatewayResponse.Code)
		return 
	}
	fmt.Println("验签结果：",gatewayResponse.SignVerifyResult)
	getCertificateResponse := efaceid.GetCertificateResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent),&getCertificateResponse)
	fmt.Println("服务器返回业务数据 : ",getCertificateResponse.ToJsonStr())
}