package main

import(
	"fmt"
	"testing"
	"time"
	

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/zfi"
	"com.es.gateway.sdk/sdk/util"
)
/**
 * H5实名认证（兼容版本） -- 此版本体检只用于浏览器，小程序，公众号，不要用在APP上

     1. 服务开通地址：https://market.aliyun.com/products/57000002/cmapi00043344.html?spm=5176.shop.result.2.22d928d8JU0OOn&innerSource=search
     2. 参考文档地址：https://esandinfo.yuque.com/books/share/9bc3e1a2-e5ca-4805-9508-89b8d55ea149?#
     3. 管理控制台地址：https://openali.esandcloud.com
     4. H5试用DEMO地址 (仅活体)：https://wxapp.dev.esandcloud.com/h5demo/esand/living
     5. 小程序使用地址：微信上搜 "一砂可信链接"，或者扫如下二维码：


     6. 域名白名单配置可参考（微信小程序）：https://kf.qq.com/touch/sappfaq/171102ue6viI171102jm63uy.html
 */
func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode :=  "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	zfiInitRequest := zfi.ZfiInitRequest{}
    zfiInitRequest.BizId = time.Now().Format("20060102150405") + "_" + util.RandStr(6)
	zfiInitRequest.IdName = "张三"
	zfiInitRequest.IdNumber = "1025591988666254299"
	zfiInitRequest.ReturnUrl = "http://www.baidu.com"
	zfiInitRequest.NotifyUrl = "http://www.baidu.com"
	zfiInitRequest.Type = "0"
	zfiInitRequest.LivingType = "1"
	zfiInitRequest.NeedVideo = "false"

	bizContent := zfiInitRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(zfi.ZfiInitRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	//解析返回的报文内容
	fmt.Println("服务器返回内容为: ",rspMsg)

}
/**
*获取认证结果
**/
func Test2(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode :=  "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	ossDeleteRequest := zfi.OssDeleteRequest{}
    ossDeleteRequest.BusinessId = "117577928278938"//初始化返回的业务id（注意切勿让中间人替换此ID）
	
	bizContent := ossDeleteRequest.ToJsonStr()
    rspMsg := gateway.SendToGateWay(zfi.GetResultRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	//解析返回的报文内容
	fmt.Println("服务器返回内容为: ",rspMsg)

}
/**
*删除缓存数据,参考文档：https://esandinfo.yuque.com/yv6e1k/aa4qsg/vcqbbv08gb60q4wu
**/
func Test3(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode :=  "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	ossDeleteRequest := zfi.OssDeleteRequest{}
    ossDeleteRequest.BusinessId = "117577928278938"//初始化返回的业务id（注意切勿让中间人替换此ID）
	
	bizContent := ossDeleteRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(zfi.OssDeleteRequestALIYUN_URL, bizContent)
	//解析返回的报文内容
	fmt.Println("服务器返回内容为: ",rspMsg)
}
