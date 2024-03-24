package zszn

import(
	"encoding/json"
)

type NameIDcardAccountVerifyResponse struct{
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
	*二级状态信息(resultMsg)描述的对应关系
	**/
	ResultCode string `json:"resultCode"`
	/**
	*结果描述
	**/
	ResultMsg string `json:"resultMsg"`
	
}
/**
 * 转换成JSON字符串
 */
 func (self NameIDcardAccountVerifyResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}