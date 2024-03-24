package mno_es2

import(
	"encoding/json"
)

type VerifyMobileRequest struct{
	/**
	*APP名称
	**/
	AppName string `json:"appName"`
	/**
	*设备型号
	**/
	DeviceModel string `json:"deviceModel"`
	/**
	*应用包名
	**/
	PackageId string `json:"packageId"`
	/**
	*运行平台
	**/
	Platform string `json:"platform"`
	/**
	手机号码(此号码与获取到的本机号码进行比对)
	**/
	PhoneNum string `json:"phoneNum"`
	/**
	*认证TOKEN（APP端返回）
	**/
	Token string `json:"token"`
}
/**
*一砂云，请求action
**/
const VerifyMobileRequestAct = "mno/es2/verifyMobile"
/**
 * 转换成JSON字符串
 */
 func (self VerifyMobileRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}