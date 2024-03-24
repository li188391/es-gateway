package main

import(
	"encoding/json"
	"fmt"
	"testing"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/non_mainland_cn2"
)

func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	idNameRequest := non_mainland_cn2.IdNameRequest{}
	idNameRequest.Nation = "CHN"
	idNameRequest.IdType = "3"
	idNameRequest.Name = "李晨"
	idNameRequest.IdNo ="H60511121"

	bizContent := idNameRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(non_mainland_cn2.IdNameRequestAct,bizContent)
	if gatewayResponse.Code !="0000"{
		t.Error("服务器返回异常 : ",gatewayResponse.Code)
		return 
	}
	fmt.Println("验签结果：",gatewayResponse.SignVerifyResult)
	idNameResponse := non_mainland_cn2.IdNameResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent),&idNameResponse)
	fmt.Println("服务器返回业务数据 : ",idNameResponse.ToJsonStr())
}