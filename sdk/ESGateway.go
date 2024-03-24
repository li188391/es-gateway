package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ESGateway struct {
	/**
	 * 网关配置的算法
	 */
	algo string
	/**
	 * 网关密钥
	 */
	key string
	/**
	 * 用户的APPCODE
	 */
	appCode string
}

func (esGateway *ESGateway) Init(algo string, key string, appCode string) {
	esGateway.key = key
	esGateway.algo = algo
	esGateway.appCode = appCode
}

/**
* @param act           业务action（参考具体业务协议文档）
* @param bizContent    业务数据，（参考具体业务协议文档）(json字符串)
* @return
 */
func (esGateway *ESGateway) SendToGateWay(act string, bizContent string) GatewayResponse {
	gatewayResponse := GatewayResponse{}
	for {
		if act == "" || bizContent == "" {
			gatewayResponse.initWithCode("-1", "参数异常")
			break
		}

		// 构造请求报文
		gatewayRequest := GatewayRequest{Act: act, AppCode: esGateway.appCode, Type: esGateway.algo, BizContent: bizContent, Timestamp: time.Now().Format("20060102150405") + "000"}
		gatewayRequest.genSign(esGateway.key)
		// 发送网络请求
		targetUrl := "https://edis.esandcloud.com/gateways"
		payload := gatewayRequest.ToJsonStr()
		fmt.Println("\n\n\n                                ----------------------------\n\n")
		fmt.Println("请求报文：" + payload)
		req, _ := http.NewRequest("POST", targetUrl, strings.NewReader(payload))
		req.Header.Add("Content-Type", "application/json")
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			gatewayResponse.initWithCode("-1", err.Error())
			break
		}

		defer response.Body.Close()
		rspStatusCode := response.Status
		if rspStatusCode != "200 OK" {
			gatewayResponse.initWithCode("-1", "服务器返回异常，状态码为："+rspStatusCode)
			break
		}

		rspBody, _ := io.ReadAll(response.Body)
		fmt.Println("响应内容 : ", string(rspBody))
		fmt.Println("\n\n                                ----------------------------\n\n\n")
		// 对数据进行验签
		json.Unmarshal(rspBody, &gatewayResponse)
		gatewayResponse.verifySign(esGateway.key)
		break // 必须写，否则进入死循环
	}

	return gatewayResponse
}
