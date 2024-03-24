package main

import(
	"fmt"
	"testing"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/zszn"

)
/**
 * 手机号在线状态，服务开通地址：https://market.aliyun.com/products/57000002/cmapi00042298.html?spm=5176.shop.result.20.6b7c28d8M6UN7M&innerSource=search#sku=yuncode3629800001
 * 管理控制台地址：https://openali.esandcloud.com/
 */

/*获取活体行为*/
 func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode := "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	mobileStatusRequest := zszn.MobileStatusRequest{}
	mobileStatusRequest.Mobile = "17512044665"
	mobileStatusRequest.NotifyUrl = "http://callback.com/mobiles_state"
    
	bizContent := mobileStatusRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(zszn.MobileStatusRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}