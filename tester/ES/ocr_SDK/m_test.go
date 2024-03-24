package main

import(
	"encoding/json"
	"fmt"
	"testing"
	

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/ocr_SDK"
)
/**
*证件OCR,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/psrxgo
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	ocrInitRequest := ocr_SDK.OcrInitRequest{}
	ocrInitRequest.OcrType = "1"

	bizContent := ocrInitRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(ocr_SDK.OcrInitRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	ocrInitResponse := ocr_SDK.OcrInitResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &ocrInitResponse)
	fmt.Println("服务器返回业务数据 : ", ocrInitResponse.ToJsonStr())
}