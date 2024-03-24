package main

import(
	"encoding/json"
	"fmt"
	"testing"
	"time"
	

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/zfi"
	"com.es.gateway.sdk/sdk/util"
)
/**
*认证初始化,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/dzpzz0
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	zfiInitRequest := zfi.ZfiInitRequest{}
    zfiInitRequest.BizId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	zfiInitRequest.IdName = "张三"
	zfiInitRequest.IdNumber = "1025591988666254299"
	zfiInitRequest.ReturnUrl = "http://www.baidu.com"
	zfiInitRequest.NotifyUrl = "http://www.baidu.com"
	zfiInitRequest.Type = "0"
	zfiInitRequest.LivingType = "1"
	zfiInitRequest.NeedVideo = "false"

	bizContent := zfiInitRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(zfi.ZfiInitRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	zfiInitResponse := zfi.ZfiInitResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &zfiInitResponse)
	fmt.Println("服务器返回业务数据 : ", zfiInitResponse.ToJsonStr())
}
/**
*获取认证结果,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/xmber7
**/
func Test2(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	getResultRequest := zfi.GetResultRequest{}
    getResultRequest.BusinessId = "117577928278938"//初始化返回的业务id（注意切勿让中间人替换此ID）
	
	bizContent := getResultRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(zfi.GetResultRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	getResultResponse := zfi.GetResultResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &getResultResponse)
	fmt.Println("服务器返回业务数据 : ", getResultResponse.ToJsonStr())
}
/**
*删除缓存数据,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/vcqbbv08gb60q4wu
**/
func Test3(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	ossDeleteRequest := zfi.OssDeleteRequest{}
    ossDeleteRequest.BusinessId = "117577928278938"//初始化返回的业务id（注意切勿让中间人替换此ID）
	
	bizContent := ossDeleteRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(zfi.OssDeleteRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	ossDeleteResponse := zfi.OssDeleteResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &ossDeleteResponse)
	fmt.Println("服务器返回业务数据 : ", ossDeleteResponse.ToJsonStr())
}
