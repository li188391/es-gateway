package main

import(
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
 * 人脸比对,服务开通地址：https://market.aliyun.com/products/57124001/cmapi00045552.html?spm=5176.shop.result.78.553928d8UHqi88&innerSource=search#sku=yuncode3955200001
 * 管理控制台地址：https://openali.esandcloud.com/
**/

func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode :=  "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
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
	rspMsg := gateway.SendToGateWay(face_v2.VerifyV2RequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	fmt.Println("服务器返回内容为: ",rspMsg)
}