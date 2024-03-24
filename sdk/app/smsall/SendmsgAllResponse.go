package smsall

import(
	"encoding/json"
)

type SendmsgAllResponse struct{
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*此次短信发送的ID
	**/
	ExternalId string `json:"externalId"`
	/**
	*此次请求ID
	**/
	RequestId string `json:"requestId"`
}
/**
 * 转换成JSON字符串
 */
 func (self SendmsgAllResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}