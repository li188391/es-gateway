package es2

import(
	"encoding/json"
)

type GetMobileResponse struct{
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
	/**
	*获取的手机号码
	**/
	PhoneNum float64 `json:"phoneNum"`
}
/**
 * 转换成JSON字符串
 */
 func (self GetMobileResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}