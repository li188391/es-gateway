package main

import(
	"encoding/json"
	"encoding/base64"
	"io/ioutil"
	"log"
	"fmt"
	"testing"
	"time"


	"com.es.gateway.sdk/sdk"
    "com.es.gateway.sdk/sdk/app/rpverifyExt"
	"com.es.gateway.sdk/sdk/util"

)

func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	uploadIDIMGRequest := rpverifyExt.UploadIDIMGRequest{}
    uploadIDIMGRequest.BizId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
    path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	uploadIDIMGRequest.IDImg = base64Str
	uploadIDIMGRequest.CountryISOCode = "US"
	uploadIDIMGRequest.IDType = "passport"
    uploadIDIMGRequest.IDNumber = "102020000000"
	uploadIDIMGRequest.Name = "PONY"
    bizContent := uploadIDIMGRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(rpverifyExt.UploadIDIMGRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	uploadIDIMGResponse := rpverifyExt.UploadIDIMGResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &uploadIDIMGResponse)
	fmt.Println("服务器返回业务数据 : ", uploadIDIMGResponse.ToJsonStr())

}

func Test2(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	//构造请求报文
	getVerifyURLRequest := rpverifyExt.GetVerifyURLRequest{}
	getVerifyURLRequest.Eid = "20231127194945330_qfeu"
	getVerifyURLRequest.NotifyUrl = "https://face.es.com/notifyUrl"
	getVerifyURLRequest.ReturnUrl = "https://face.es.com/returnUrl"
	getVerifyURLRequest.LivingType = "13"
    
	bizContent := getVerifyURLRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(rpverifyExt.GetVerifyURLRequestAct,bizContent)
	if gatewayResponse.Code !="0000"{
		t.Error("服务器返回异常 : ",gatewayResponse.Code)
		return 
	}
	fmt.Println("验签结果：",gatewayResponse.SignVerifyResult)
	getVerifyURLResponse := rpverifyExt.GetVerifyURLResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent),&getVerifyURLResponse)
	fmt.Println("服务器返回业务数据 : ",getVerifyURLResponse.ToJsonStr())
}