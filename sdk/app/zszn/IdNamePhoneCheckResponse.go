package zszn

import(
	"encoding/json"
)

type IdNamePhoneCheckResponse struct{
	/**
	*请求ID，用于关联请求日志
	**/
	RequestId string `json:"requestId"`
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*响应描述
	**/
	Msg string `json:"msg"`
	/**
	*校验结果：1：认证信息匹配，2：认证信息不匹配，3：无法验证，-1：异常情况
	**/
	Result string `json:"result"`
	/**
	*结果描述
	**/
	ResultMsg string `json:"resultMsg"`
	/**
	*运营商名称
	**/
	Isp string `json:"isp"`
}
/**
 * 转换成JSON字符串
 */
 func (self IdNamePhoneCheckResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}