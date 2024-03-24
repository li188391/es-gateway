package main

import(
	"fmt"
	"testing"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/zszn"

)
 /**
   * 运营商三要素校验，参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/phim68
   * 服务开通地址：https://market.aliyun.com/products/57000002/cmapi00042296.html?spm=5176.shop.result.23.6b7c28d860HlR0&innerSource=search
 **/

 func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode := "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	idNamePhoneCheckRequest := zszn.IdNamePhoneCheckRequest{}
	idNamePhoneCheckRequest.IdName = "李翠花"
	idNamePhoneCheckRequest.IdNumber = "596223555698546855"
	idNamePhoneCheckRequest.Mobile = "16599998888"
    
	bizContent := idNamePhoneCheckRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(zszn.IdNamePhoneCheckRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}