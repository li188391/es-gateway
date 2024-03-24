package main

import(
	"encoding/json"
	"fmt"
	"testing"
	"encoding/base64"
	"io/ioutil"
	"log"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/non_mainland_cn3"
)

func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	idNameImageRequest := non_mainland_cn3.IdNameImageRequest{}
	idNameImageRequest.Nation = "CHN"
	idNameImageRequest.IdType = "1"
	idNameImageRequest.Name = "黄小橘"
	idNameImageRequest.IdNO = "751114215478"
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	idNameImageRequest.PersonImg = base64Str

	bizContent := idNameImageRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(non_mainland_cn3.IdNameImageRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	idNameImageResponse := non_mainland_cn3.IdNameImageResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &idNameImageResponse)
	fmt.Println("服务器返回业务数据 : ", idNameImageResponse.ToJsonStr())
}