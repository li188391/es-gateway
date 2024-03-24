package main

import(
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/non_mainland_cnH5"
	"com.es.gateway.sdk/sdk/util"
)
/**
*获取认证链接,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/pgvrn6ignukl8l5n
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
    getAuthUrlRequest := non_mainland_cnH5.GetAuthUrlRequest{}
	getAuthUrlRequest.BizId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	getAuthUrlRequest.LivingType = "13"
	getAuthUrlRequest.IdType = "1"
	getAuthUrlRequest.CertName = "张三"
	getAuthUrlRequest.CertNo = "1025591988666254299"
	getAuthUrlRequest.ReturnUrl = "http://www.baidu.com"
	getAuthUrlRequest.NotifyUrl = "http://www.baidu.com"

	bizContent := getAuthUrlRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(non_mainland_cnH5.GetAuthUrlRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	getAuthUrlResponse := non_mainland_cnH5.GetAuthUrlResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &getAuthUrlResponse)
	fmt.Println("服务器返回业务数据 : ", getAuthUrlResponse.ToJsonStr())
}