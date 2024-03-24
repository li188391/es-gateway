package main

import(
	"fmt"
	"testing"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/zszn"

)
 /**
  * 服务开通地址：https://market.aliyun.com/products/57000002/cmapi00042295.html?spm=5176.shop.result.23.6b7c28d89pqdoz&innerSource=search#sku=yuncode3629500001
  *银行卡实名三要素，参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/iftlbl
 **/

 func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode := "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	nameIDCardAccountVerifyRequest := zszn.NameIDCardAccountVerifyRequest{}
	nameIDCardAccountVerifyRequest.AccountNo = "62008856445623558"
	nameIDCardAccountVerifyRequest.IdName = "李翠花"
	nameIDCardAccountVerifyRequest.IdNumber = "596223555698546855"
    
	bizContent := nameIDCardAccountVerifyRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(zszn.NameIDCardAccountVerifyRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}
 /**
 * 银行卡四要素，参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/guals8
 **/
 func Test2(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode := "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	nameIDCardPhoneAccountVerifyRequest := zszn.NameIDCardPhoneAccountVerifyRequest{}
	nameIDCardPhoneAccountVerifyRequest.AccountNo = "62008856445623558"
	nameIDCardPhoneAccountVerifyRequest.IdName = "李翠花"
	nameIDCardPhoneAccountVerifyRequest.IdNumber = "596223555698546855"
	nameIDCardPhoneAccountVerifyRequest.Mobile = "16599998888"

    
	bizContent := nameIDCardPhoneAccountVerifyRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(zszn.NameIDCardPhoneAccountVerifyRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}