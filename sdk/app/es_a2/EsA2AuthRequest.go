package es_a2

import(
	"encoding/json"
)

type EsA2AuthRequest struct{
	/**
	*业务ID(需保证唯一，可选)
	**/
	TransId string `json:"transId"`
	/**
	*身份证号码
	**/
	IdNumber string `json:"idNumber"`
	/**
	*姓名
	**/
	Name string `json:"name"`

}
/**
*一砂云，请求action
**/
const EsA2AuthRequestAct = "ctid/es_a2/auth"
/**
 * 转换成JSON字符串
 */
 func (self EsA2AuthRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
