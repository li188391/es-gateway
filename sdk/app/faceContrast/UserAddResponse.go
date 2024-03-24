package faceContrast

import(
	"encoding/json"
)

type UserAddResponse struct{
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
	*执行add或者update
	**/
	Action string `json:"action"`
}
/**
 * 转换成JSON字符串
 */
 func (self UserAddResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}