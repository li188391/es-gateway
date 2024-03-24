package es_a3

import(
	"encoding/json"
)

type EsA3AuthResponse struct{
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*人像比对分数，一砂默认阈值为750分，取值范围为0-1000
	**/
	Score string `json:"score"`
	/**
	*业务ID
	**/
	TransId string `json:"transId"`
}
/**
 * 转换成JSON字符串
 */
 func (self EsA3AuthResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}