package main

import(
	"fmt"
	"testing"
	"encoding/base64"
	"io/ioutil"
	"log"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/edis_ctid3"

	
)
/**
 * 实名三要素，产品开通地址：https://market.aliyun.com/products/57000002/cmapi00037640.html?spm=5176.shop.result.17.6b7c28d89pqdoz&innerSource=search
 * 文档地址：https://esandinfo.yuque.com/books/share/f55eecd0-c7d1-4625-8245-67aa9f9b89a1?#
 * 管理控制台地址：https://openali.esandcloud.com/
 */

func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode :=  "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	idNameImageRequest := edis_ctid3.IdNameImageRequest{}
	idNameImageRequest.IdCard = "110101199003078312"
	idNameImageRequest.Name = "张三"
	path :="resources/cn_idcard.jpeg"	
	imageFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(imageFile)
	idNameImageRequest.Image = base64Str
	idNameImageRequest.HadLiving = "true";
	idNameImageRequest.SecLevel = "loose"
	

	bizContent := idNameImageRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(edis_ctid3.IdNameImageRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	fmt.Println("服务器返回内容为: ",rspMsg)
}