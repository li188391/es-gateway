package main

import(
	"fmt"
	"testing"
	"time"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/h5mno"
	"com.es.gateway.sdk/sdk/util"

)
/**
 *   参考文档：https://esandinfo.yuque.com/books/share/36b5f381-12de-4fbd-8b93-429653ca36e9?#
 *   服务器开通地址：https://market.aliyun.com/products/57126001/cmapi00046908.html?spm=5176.shop.result.20.6b7c28d87JsmuD&innerSource=search#sku=yuncode4090800001
 *   管理控制台地址：https://openali.esandcloud.com
 **/

/*初始化认证*/
 func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode := "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	h5mnoInitRequest := h5mno.H5mnoInitRequest{}
	h5mnoInitRequest.BizId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	h5mnoInitRequest.PhoneNumber = "17512044665"
	h5mnoInitRequest.ReturnUrl = "http://www.baidu.com"
    
	bizContent := h5mnoInitRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(h5mno.H5mnoInitRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}