package zfi

import(
	"encoding/json"
)

type OssDeleteResponse struct{
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*日志ID
	**/
	LogID string `json:"logID"`
	/**
	*是否计费，此接口不计费，固定为0
	**/
	BillingResult string `json:"billingResult"`
}
/**
 * 转换成JSON字符串
 */
 func (self OssDeleteResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}