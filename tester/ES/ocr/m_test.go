package main

import(
	"encoding/json"
	"fmt"
	"testing"
	"encoding/base64"
	"io/ioutil"
	"log"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/ocr"
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
	ocrVerifyRequest := ocr.OcrVerifyRequest{}
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	ocrVerifyRequest.File = base64Str
	ocrVerifyRequest.OcrType = "1"

	bizContent := ocrVerifyRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(ocr.OcrVerifyRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	ocrVerifyResponse := ocr.OcrVerifyResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &ocrVerifyResponse)
	fmt.Println("服务器返回业务数据 : ", ocrVerifyResponse.ToJsonStr())
}