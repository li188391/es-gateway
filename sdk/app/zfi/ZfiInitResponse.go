package zfi

import(
	"encoding/json"
)

type ZfiInitResponse struct{
	/**
	*初始化业务Id,结果查证时需要
	**/
	BusinessId string `json:"businessId"`
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*此次请求ID
	**/
	RequestId string `json:"requestId"`
	/**
	*认证链接
	**/
	Url string `json:"url"`
}
/**
 * 转换成JSON字符串
 */
 func (self ZfiInitResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}