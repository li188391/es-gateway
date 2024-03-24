package main

import(
	"encoding/json"
	"fmt"
	"testing"

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/livingdetect"
)
/**
*获取TOKEN,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/chz53t
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	livingdetectRequest := livingdetect.LivingdetectRequest{}
	livingdetectRequest.InitMsg = "@Android,2,8,1658484566@NibhdDZK9ja0gwaTLRo5KpckYi6y6tRCjWFqAMZZzNNeLSsfNtVddN8/dzI6HJFUp+tg1c0Qc1DqpP/CDriJlOKEgcGvlXLUP3fhv2wYGwT0SDRAv1x2obTbWn6N0AVhHkUApE1GjEZlSVHPQn350vSUa8a/Hm4RBW7D0Uu8BjcR6FS8Q/7j2klpH0lRTmsZI7oxUA2qAfX7pZ7e068ZR613bdSzwbkSvaer7sSqmJnMwvp0f3oRZ5ZC0NrUuK2WPkETzzklCbUhE24yB1e3xN11HG52HX+/MfyHuzPGoLGNd4VfPzfaul+m6hfl3Aaoy8VE9tjk7ewdaM8TtsDN4BMvcAJ75CKjfxCdkXOqdq2XVXtmMnGNjse0DpZfzhFko3G1HtOCPliU+b5ytIB0DNi/aA2ec0Ojgqu0JUjkilxq52jIns05Cj6eDJ+MeOWc1pauxmBH8qkKLPK7VCQxpYRn8PphMLcv6+w6Nc92rEjXGN4gL6bndjrnMEpPeK/eUf+Oq9TfYQ7YyBWOIPse8GJTNcZocxR3j8KRHeO70j0dDme+DJqDblcz2yU2AzgV1Gm6VR8rDw3fN1KkMVgc9w=="
	
	bizContent := livingdetectRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(livingdetect.LivingdetectRequestAct,bizContent)
	if gatewayResponse.Code !="0000"{
		t.Error("服务器返回异常 : ",gatewayResponse.Code)
		return 
	}
	fmt.Println("验签结果：",gatewayResponse.SignVerifyResult)
	livingdetectResponse := livingdetect.LivingfdetectResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent),&livingdetectResponse)
	fmt.Println("服务器返回业务数据 : ",livingdetectResponse.ToJsonStr())
}
/**
*获取活体检测结果,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/rp6dws
**/
func Test2(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	livingdetectVerifyRequest := livingdetect.LivingdetectVerifyRequest{}
    livingdetectVerifyRequest.Token = "kbi+dIKhrfdn177jS3yYw9Rv+GfkJXNbIRpJUbqSXRbqbw+HtkFAv4JT9doMBWZXRsTmFoW+9VJx+yIs/Hg35ZNwVKyXE0zGyQtL6RBxa0E="
	livingdetectVerifyRequest.VerifyMsg = "NULL"
	
	bizContent := livingdetectVerifyRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(livingdetect.LivingdetectVerifyRequestAct,bizContent)
	if gatewayResponse.Code !="0000"{
		t.Error("服务器返回异常 : ",gatewayResponse.Code)
		return 
	}
	fmt.Println("验签结果：",gatewayResponse.SignVerifyResult)
	livingdetectVerifyResponse := livingdetect.LivingdetectVerifyResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent),&livingdetectVerifyResponse)
	fmt.Println("服务器返回业务数据 : ",livingdetectVerifyResponse.ToJsonStr())
}