package es2

import(
	"encoding/json"
)

type GetMobileRequest struct{
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
	*认证TOKEN（APP端返回）
	**/
	Token string `json:"token"`
}
/**
*一砂云，请求action
**/
const GetMobileRequestAct = "mno/es2/getMobile"
/**
 * 转换成JSON字符串
 */
 func (self GetMobileRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}