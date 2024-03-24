package main

import(
	"fmt"
	"testing"
	
	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/bank"

)
 /**
 * 银行卡信息查询 服务开通地址：https://market.aliyun.com/products/57000002/cmapi00052920.html?spm=5176.shop.result.14.6b7c28d8EEIsjj&innerSource=search
 * 参考文档：https://esandinfo.yuque.com/books/share/8bbcd12a-d91f-43a2-8568-89dcfe348045?#
 * 管理控制台地址： https://openali.esandcloud.com/
 */

 func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode := "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	bankCardInfoRequest := bank.BankCardInfoRequest{}
	bankCardInfoRequest.AccountNo = "62008856445623558"
	
    
	bizContent := bankCardInfoRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(bank.BankCardInfoRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	
	fmt.Println("服务器返回内容为: ",rspMsg)
}