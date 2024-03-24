package main

import(
	"encoding/json"
	"fmt"
	"testing"
	"time"
	

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/non_mainland_cn"
	"com.es.gateway.sdk/sdk/util"
)

func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
    nonMainlandCnRequest := non_mainland_cn.NonMainlandCnRequest{}
	nonMainlandCnRequest.BizId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	nonMainlandCnRequest.InitMsg = "@Android,2,8,1658484566@NibhdDZK9ja0gwaTLRo5KpckYi6y6tRCjWFqAMZZzNNeLSsfNtVddN8/dzI6HJFUp+tg1c0Qc1DqpP/CDriJlOKEgcGvlXLUP3fhv2wYGwT0SDRAv1x2obTbWn6N0AVhHkUApE1GjEZlSVHPQn350vSUa8a/Hm4RBW7D0Uu8BjcR6FS8Q/7j2klpH0lRTmsZI7oxUA2qAfX7pZ7e068ZR613bdSzwbkSvaer7sSqmJnMwvp0f3oRZ5ZC0NrUuK2WPkETzzklCbUhE24yB1e3xN11HG52HX+/MfyHuzPGoLGNd4VfPzfaul+m6hfl3Aaoy8VE9tjk7ewdaM8TtsDN4BMvcAJ75CKjfxCdkXOqdq2XVXtmMnGNjse0DpZfzhFko3G1HtOCPliU+b5ytIB0DNi/aA2ec0Ojgqu0JUjkilxq52jIns05Cj6eDJ+MeOWc1pauxmBH8qkKLPK7VCQxpYRn8PphMLcv6+w6Nc92rEjXGN4gL6bndjrnMEpPeK/eUf+Oq9TfYQ7YyBWOIPse8GJTNcZocxR3j8KRHeO70j0dDme+DJqDblcz2yU2AzgV1Gm6VR8rDw3fN1KkMVgc9w=="
	nonMainlandCnRequest.Nation = "CHN"
	nonMainlandCnRequest.IdType = "1"
	nonMainlandCnRequest.CertName = "张三"
	nonMainlandCnRequest.CertNo = "H60511121"

	bizContent := nonMainlandCnRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(non_mainland_cn.NonMainlandCnRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	nonMainlandCnResponse := non_mainland_cn.NonMainlandCnResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &nonMainlandCnResponse)
	fmt.Println("服务器返回业务数据 : ", nonMainlandCnResponse.ToJsonStr())
}
/**
* 获取认证授权认证结果, 参考文档: https://esandinfo.yuque.com/yv6e1k/aa4qsg/igwq5xfffuk9r96t
**/
func Test2(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
    NonMainlandCnVerifyRequest := non_mainland_cn.NonMainlandCnVerifyRequest{}
	NonMainlandCnVerifyRequest.Token = "kbi+dIKhrfdn177jS3yYw9Rv+GfkJXNbIRpJUbqSXRbqbw+HtkFAv4JT9doMBWZXkaQLxBc44BedMtKJs/d76XARaVHwzvjoX/uVbiOqBJ4="
	NonMainlandCnVerifyRequest.VerifyMsg = "NULL"

	bizContent := NonMainlandCnVerifyRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(non_mainland_cn.NonMainlandCnVerifyRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	NonMainlandCnVerifyResponse := non_mainland_cn.NonMainlandCnVerifyResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &NonMainlandCnVerifyResponse)
	fmt.Println("服务器返回业务数据 : ", NonMainlandCnVerifyResponse.ToJsonStr())
}
