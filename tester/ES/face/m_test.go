package main

import(
	"encoding/json"
	"fmt"
	"testing"
	"encoding/base64"
	"io/ioutil"
	"log"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/face"
	
)
/**
*创建人脸数据库,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/wxkbyy
**/
func Test1(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	createFaceDbRequest := face.CreateFaceDbRequest{}
    createFaceDbRequest.DbName = "esandinfo"
	createFaceDbRequest.Description = "一砂员工人脸数据库"

	bizContent := createFaceDbRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(face.CreateFaceDbRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	createFaceDbResponse := face.CreateFaceDbResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &createFaceDbResponse)
	fmt.Println("服务器返回业务数据 : ", createFaceDbResponse.ToJsonStr())
}
/**
*增加人脸样本,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/vigu0t
**/
func Test2(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	addFaceRequest := face.AddFaceRequest{}
    addFaceRequest.DbName = "esandinfo"
    addFaceRequest.EntityId = "reidlee"
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	addFaceRequest.Image = base64Str
	addFaceRequest.ExtraData = "rd0010"

	bizContent := addFaceRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(face.AddFaceRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	addFaceResponse := face.AddFaceResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &addFaceResponse)
	fmt.Println("服务器返回业务数据 : ", addFaceResponse.ToJsonStr())
}
/**
*更新人脸样本,参考文档:https://esandinfo.yuque.com/yv6e1k/aa4qsg/azmg27
**/
func Test3(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	updateFaceRequest := face.UpdateFaceRequest{}
    updateFaceRequest.DbName = "esandinfo"
    updateFaceRequest.EntityId = "reidlee"
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	updateFaceRequest.Image = base64Str
	updateFaceRequest.ExtraData = "rd0010"

	bizContent := updateFaceRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(face.UpdateFaceRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	updateFaceResponse := face.UpdateFaceResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &updateFaceResponse)
	fmt.Println("服务器返回业务数据 : ", updateFaceResponse.ToJsonStr())
}
/**
*人脸搜索,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/gheu8p
**/
func Test4(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	searchFaceRequest := face.SearchFaceRequest{}
    searchFaceRequest.DbName = "esandinfo"
    searchFaceRequest.EnableLiving = true
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	searchFaceRequest.Image = base64Str


	bizContent := searchFaceRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(face.SearchFaceRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	searchFaceResponse := face.SearchFaceResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &searchFaceResponse)
	fmt.Println("服务器返回业务数据 : ", searchFaceResponse.ToJsonStr())
}
/**
*删除人脸样本，参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/kcm58z
**/
func Test5(t *testing.T){
	appCode := "d2808c1338ce01f3e3efdb486f9effb9"
	algo := "3"
	key := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhdQdw6uTGz+X4e+/eGpu4dcKrfrjQAe0tTQyn6TuKBobcsUsESkx9jcCeG4cpuxmN8jNeDAZsbTujNLD/aYl2g=="
	esGateway := sdk.ESGateway{}
	esGateway.Init(algo, key, appCode)
	// 构造请求报文
	deleteFaceRequest := face.DeleteFaceRequest{}
    deleteFaceRequest.DbName = "esandinfo"
    deleteFaceRequest.EntityId = "reidlee"
	

	bizContent := deleteFaceRequest.ToJsonStr()
	gatewayResponse := esGateway.SendToGateWay(face.DeleteFaceRequestAct, bizContent)
	if gatewayResponse.Code != "0000" {
		t.Error("服务返回异常 : ", gatewayResponse.Code)
		return
	}

	fmt.Println("验签结果 : ", gatewayResponse.SignVerifyResult)
	deleteFaceResponse := face.DeleteFaceResponse{}
	json.Unmarshal([]byte(gatewayResponse.BizContent), &deleteFaceResponse)
	fmt.Println("服务器返回业务数据 : ", deleteFaceResponse.ToJsonStr())
}

