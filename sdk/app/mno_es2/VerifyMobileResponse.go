package mno_es2

import(
	"encoding/json"
)

type VerifyMobileResponse struct{
	/**
	*业务号
	**/
	TransId string `json:"transId"`
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*响应描述
	**/
	Msg string `json:"msg"`
}
/**
 * 转换成JSON字符串
 */
 func (self VerifyMobileResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}