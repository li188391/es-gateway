package faceContrast

import(
	"encoding/json"
)

type UserDeleteResponse struct{
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
}
/**
 * 转换成JSON字符串
 */
 func (self UserDeleteResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}