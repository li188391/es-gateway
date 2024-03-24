package main

import(
	"fmt"
	"testing"
	"encoding/base64"
	"io/ioutil"
	"log"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/zszn"

	
)
/**
 * 图片活体检测 产品地址：https://market.aliyun.com/products/57000002/cmapi00043301.html?spm=5176.shop.result.35.6b7c28d8D23Kzv&innerSource=search#sku=yuncode3730100001
 * 管理控制台地址： https://openali.esandcloud.com
**/

func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode :=  "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	sLivingVerifyRequest := zszn.SLivingVerifyRequest{}
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	sLivingVerifyRequest.Image = base64Str
	

	bizContent := sLivingVerifyRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(zszn.SLivingVerifyRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	fmt.Println("服务器返回内容为: ",rspMsg)
}