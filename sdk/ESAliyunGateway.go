package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/li188391/es-gateway/sdk/util"
)

type ESAlyunGateway struct {
	/**
	 * 实例化
	 * @param appCode 获取 appCode,参考 : https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	 */
	appCode string
}

func (eSAlyunGateway *ESAlyunGateway) Init(appCode string) {
	eSAlyunGateway.appCode = appCode
}

/**
* @param url         阿里云的URL地址
* @param bizContent    业务数据，（参考具体业务协议文档）(json字符串)
* @return
 */
func (eSAlyunGateway *ESAlyunGateway) SendToGateWay(url string, bizContent string) string {
	gatewayResponse := GatewayResponse{}
	rspBody := ""
	for {
		if url == "" || bizContent == "" {
			gatewayResponse.initWithCode("-1", "参数异常")
			break
		}
		body := make(map[string]string)
		err := json.Unmarshal([]byte(bizContent), &body)
		if err != nil {
			gatewayResponse.initWithCode("-1", err.Error())
			break
		}
		jsonString, err := json.Marshal(body)
		if err != nil {
			gatewayResponse.initWithCode("-1", err.Error())
			break
		}
		fmt.Println("请求报文：" + string(jsonString))
		formData := ""
		for key, value := range body {
			formData += key + "=" + value + "&"
		}
		formData = strings.TrimSuffix(formData, "&")
		reqBody := strings.NewReader(formData)
		req, _ := http.NewRequest("POST", url, reqBody)
		req.Header.Add("Authorization", "APPCODE "+eSAlyunGateway.appCode)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.Header.Add("X-Ca-Nonce", util.RandStr(16))
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			gatewayResponse.initWithCode("-1", err.Error())
			break
		}

		defer response.Body.Close()
		rspStatusCode := response.Status
		if rspStatusCode != "200 OK" {
			gatewayResponse.initWithCode("-1", "服务器返回异常，状态码为："+rspStatusCode)
		}

		respBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			gatewayResponse.initWithCode("-1", "读取响应体失败")
			break
		}
		rspBody = string(respBody)
		fmt.Println("body内容 : ", rspBody)
		fmt.Println("Header内容 : ", response.Header)
		break
	}
	return rspBody

}
