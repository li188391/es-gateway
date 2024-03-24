package zszn

import(
	"encoding/json"
)

type MobileStatusResponse struct{
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*响应描述
	**/
	Msg string `json:"msg"`
	/**
	*状态值
	**/
	Result string `json:"result"`
	/**
	*结果描述
	**/
	ResultMsg string `json:"resultMsg"`
	/**
	*手机号平台
	**/
	Isp string `json:"isp"`
	/**
	*请求ID
	**/
	RequestId string `json:"requestId"`

}
/**
 * 转换成JSON字符串
 */
 func (self MobileStatusResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}