package faceContrast

import(
	"encoding/json"
)

type InitResponse struct{
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*请求ID
	**/
	RequestId string `json:"requestId"`
	/**
	*认证token
	**/
	Token string `json:"token"`
}
/**
 * 转换成JSON字符串
 */
 func (self InitResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}