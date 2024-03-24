package zszn

import(
	"encoding/json"
)

type NameIDCardPhoneAccountVerifyRequest struct{
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
	/**
	*银行卡绑定的预留手机号码
	**/
	Mobile string `json:"mobile"`
}
/**
*一砂云，请求action
**/
const NameIDCardPhoneAccountVerifyRequestAct = "comms/zszn/nameIDCardPhoneAccountVerify"
/**
*阿里云请求URL
**/
const NameIDCardPhoneAccountVerifyRequestALIYUN_URL = "http://ediscard.market.alicloudapi.com/comms/zszn/nameIDCardPhoneAccountVerify"
/**
 * 转换成JSON字符串
 */
 func (self NameIDCardPhoneAccountVerifyRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}