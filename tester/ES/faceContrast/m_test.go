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
	"com.es.gateway.sdk/sdk/app/faceContrast"
	"com.es.gateway.sdk/sdk/util"
)
/**
*用户注册，参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/iegzg3
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
    userAddRequest := faceContrast.UserAddRequest{}
	userAddRequest.Uuid = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	userAddRequest.Image = base64Str
	userAddRequest.VerifyMsg = "asdasldhjaslkdaskldasjdasdasd"
	userAddRequest.Token = "asjkldzxmcvnaskldnasldkasjd"

	bizContent := userAddRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(faceContrast.UserAddRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	userAddResponse := faceContrast.UserAddResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &userAddResponse)
	fmt.Println("服务器返回业务数据 : ", userAddResponse.ToJsonStr())
}
/**
*删除用户信息,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/al70xn
**/
func Test2(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
    userDeleteRequest := faceContrast.UserDeleteRequest{}
	userDeleteRequest.Uuid = "1231234314"
	

	bizContent := userDeleteRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(faceContrast.UserDeleteRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	userDeleteResponse := faceContrast.UserDeleteResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &userDeleteResponse)
	fmt.Println("服务器返回业务数据 : ", userDeleteResponse.ToJsonStr())
}
/**
*获取认证TOKEN,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/auvzrg
*/
func Test3(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
    initRequest := faceContrast.InitRequest{}
	initRequest.BizId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	initRequest.InitMsg = "@Android,2,8,1658484566@NibhdDZK9ja0gwaTLRo5KpckYi6y6tRCjWFqAMZZzNNeLSsfNtVddN8/dzI6HJFUp+tg1c0Qc1DqpP/CDriJlOKEgcGvlXLUP3fhv2wYGwT0SDRAv1x2obTbWn6N0AVhHkUApE1GjEZlSVHPQn350vSUa8a/Hm4RBW7D0Uu8BjcR6FS8Q/7j2klpH0lRTmsZI7oxUA2qAfX7pZ7e068ZR613bdSzwbkSvaer7sSqmJnMwvp0f3oRZ5ZC0NrUuK2WPkETzzklCbUhE24yB1e3xN11HG52HX+/MfyHuzPGoLGNd4VfPzfaul+m6hfl3Aaoy8VE9tjk7ewdaM8TtsDN4BMvcAJ75CKjfxCdkXOqdq2XVXtmMnGNjse0DpZfzhFko3G1HtOCPliU+b5ytIB0DNi/aA2ec0Ojgqu0JUjkilxq52jIns05Cj6eDJ+MeOWc1pauxmBH8qkKLPK7VCQxpYRn8PphMLcv6+w6Nc92rEjXGN4gL6bndjrnMEpPeK/eUf+Oq9TfYQ7YyBWOIPse8GJTNcZocxR3j8KRHeO70j0dDme+DJqDblcz2yU2AzgV1Gm6VR8rDw3fN1KkMVgc9w=="

	bizContent := initRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(faceContrast.InitRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	initResponse := faceContrast.InitResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &initResponse)
	fmt.Println("服务器返回业务数据 : ", initResponse.ToJsonStr())
}
/**
*获取人脸比对结果,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/lq1hch
**/
func Test4(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
    verifyRequest := faceContrast.VerifyRequest{}
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	verifyRequest.Image = base64Str
	verifyRequest.Uuid = "1231234314"
	verifyRequest.ImageUrl = "http://www.baidu.com"
	verifyRequest.Token ="kbi+dIKhrfdn177jS3yYw9Rv+GfkJXNbIRpJUbqSXRbqbw+HtkFAv4JT9doMBWZXjJ+QuLkIXnA2ogFG9zrPny09nWvWP3n/6oJblI8Be/E="
	verifyRequest.VerifyMsg = "asdjhasjkzxclkasldkjasldasdasd"

	bizContent := verifyRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(faceContrast.VerifyRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	verifyResponse := faceContrast.VerifyResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &verifyResponse)
	fmt.Println("服务器返回业务数据 : ", verifyResponse.ToJsonStr())
}
