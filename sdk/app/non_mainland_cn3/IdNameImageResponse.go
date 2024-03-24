package non_mainland_cn3

import(
	"encoding/json"
)

type IdNameImageResponse struct{
	/**
	*请求id
	**/
	RequestId string `json:"requestId"`
	/**
	*响应状态码
	**/
	Code string `json:"code"`
	/**
	*相应描述
	**/
	Msg string `json:"msg"`
	/**
	*是否收费：true:收费，false:未收费
	**/
	Charge bool `json:"charge"`
}
/**
 * 转换成JSON字符串
 */
 func (self IdNameImageResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}