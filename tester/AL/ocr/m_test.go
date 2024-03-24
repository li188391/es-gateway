package main

import(
	"fmt"
	"testing"
	"encoding/base64"
	"io/ioutil"
	"log"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/ocr"
	
)
/**
 *   参考文档：https://esandinfo.yuque.com/books/share/36b5f381-12de-4fbd-8b93-429653ca36e9?#
 *   服务器开通地址：https://market.aliyun.com/products/57126001/cmapi00046908.html?spm=5176.shop.result.20.6b7c28d87JsmuD&innerSource=search#sku=yuncode4090800001
 *   管理控制台地址：https://openali.esandcloud.com
 **/

/*初始化认证*/
 func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode := "78ab7c4f61104339bedd7273579568ff"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	ocrVerifyRequest := ocr.OcrVerifyRequest{}
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	ocrVerifyRequest.File = base64Str
	ocrVerifyRequest.OcrType = "3"
    
	bizContent := ocrVerifyRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(ocr.OcrVerifyRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}