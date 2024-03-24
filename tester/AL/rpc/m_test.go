package main

import(
	"fmt"
	"testing"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/rpc"

)
/**
 * 实名认证（视频活体），服务开通地址：https://market.aliyun.com/products/57000002/cmapi00037639.html?spm=5176.shop.result.2.6b7c28d8M6UN7M&innerSource=search#sku=yuncode31639000012
 * 管理控制台地址：https://openali.esandcloud.com/
 */

/*获取活体行为*/
 func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode := "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	getBehaviorRequest := rpc.GetBehaviorRequest{}
    
	bizContent := getBehaviorRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(rpc.GetBehaviorRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}
/**
 * 实人认证
*/
func Test2(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode := "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	verifyRequest := rpc.VerifyRequest{}
	verifyRequest.CertName = "张三"
	verifyRequest.CertNo = "110101199003070417"
	verifyRequest.BehaviorToken = "14XWiaVF"
	verifyRequest.Video = "video"
    
	bizContent := verifyRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(rpc.VerifyRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}