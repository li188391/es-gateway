package zszn

import(
	"encoding/json"
)

type SLivingVerifyResponse struct{
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*响应描述
	**/
	Msg string `json:"msg"`
	/**
	*检测结果
	**/
	Result string `json:"result"`
	/**
	*结果描述
	**/
	ResultMsg string `json:"resultMsg"`
	/**
	*请求id
	**/
	RequestId string `json:"requestId"`
}
/**
 * 转换成JSON字符串
 */
 func (self SLivingVerifyResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}