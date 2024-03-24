package zszn

import(
	"encoding/json"
)

type NameIDCardAccountVerifyRequest struct{
	/**
	*银行卡号
	**/
	AccountNo string `json:"accountNo"`
	/**
	*姓名
	**/
	IdName string `json:"idName"`
	/**
	*身份证号
	**/
	IdNumber string `json:"idNumber"`
	
}
/**
*一砂云，请求action
**/
const NameIDCardAccountVerifyRequestAct = "comms/zszn/nameIDCardAccountVerify"
/**
*阿里云请求URL
**/
const NameIDCardAccountVerifyRequestALIYUN_URL = "http://ediscard.market.alicloudapi.com/comms/zszn/nameIDCardAccountVerify"
/**
 * 转换成JSON字符串
 */
 func (self NameIDCardAccountVerifyRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}