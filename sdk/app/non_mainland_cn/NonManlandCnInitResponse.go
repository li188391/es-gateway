package non_mainland_cn

import(
	"encoding/json"
)

type  NonMainlandCnResponse struct{
	/**
	*请求id
	**/
	RequestId string `json:"requestId"`
	/**
	*错误码
	**/
	Code string `json:"code"`
	/**
	*相应描述
	**/
	Msg string `json:"msg"`
	/**
	*认证token
	**/
	Token string `json:"token"`
}
/**
 * 转换成JSON字符串
 */
 func (self NonMainlandCnResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}