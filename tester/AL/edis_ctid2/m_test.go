package main

import(
	"fmt"
	"testing"
	

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/edis_ctid2"

)
/**
 * 实名二要素，服务购买链接：https://market.aliyun.com/products/57000002/cmapi00037641.html?spm=5176.shop.result.11.6b7c28d8sKgX9l&innerSource=search#sku=yuncode3164100001
 * 管理控制台地址：https://openali.esandcloud.com/
 */

 func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode :=  "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	idNameRequest := edis_ctid2.IdNameRequest{}
    idNameRequest.Name = "张三"
	idNameRequest.IdCard = "110101199003078312"
	idNameRequest.IdcardType = "CN"

	bizContent := idNameRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(edis_ctid2.IdNameRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}