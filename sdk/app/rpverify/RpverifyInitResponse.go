package rpverify

import(
	"encoding/json"
)
type RpverifyInitResponse struct{
	/**
	*错误码
	**/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*请求id
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
 func (self RpverifyInitResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}