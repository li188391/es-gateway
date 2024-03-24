package zszn

import(
	"encoding/json"
)

type IdNamePhoneCheckRequest struct{
	/**
	*姓名
	**/
	IdName string `json:"idName"`
	/**
	*身份证号
	**/
	IdNumber string `json:"idNumber"`
	/**
	*手机号
	**/
	Mobile string `json:"mobile"`
	
}
/**
*一砂云，请求action
**/
const IdNamePhoneCheckRequestAct = "comms/zszn/idNamePhoneCheck"
/**
*阿里云请求URL
**/
const IdNamePhoneCheckRequestALIYUN_URL = "https://edismobile.market.alicloudapi.com/comms/zszn/idNamePhoneCheck"
/**
 * 转换成JSON字符串
 */
 func (self IdNamePhoneCheckRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}