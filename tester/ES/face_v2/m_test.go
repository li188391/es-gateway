package main

import(
	"encoding/json"
	"fmt"
	"testing"
	"time"
	"encoding/base64"
	"io/ioutil"
	"log"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/face_v2"
	"com.es.gateway.sdk/sdk/util"
	
)
/**
*人脸识别(1:1比对),参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/du981s76qbvp07g5
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	verifyV2Request := face_v2.VerifyV2Request{}
    verifyV2Request.BizId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	verifyV2Request.SourceImage = base64Str
	verifyV2Request.VerifyImage = base64Str

	bizContent := verifyV2Request.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(face_v2.VerifyV2RequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	verifyV2Response := face_v2.VerifyV2Response{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &verifyV2Response)
	fmt.Println("服务器返回业务数据 : ", verifyV2Response.ToJsonStr())
}