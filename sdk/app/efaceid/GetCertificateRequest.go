package efaceid

import(
	"encoding/json"
)

type GetCertificateRequest struct{
	/**
	*APP名称
	**/
	AppName string `json:"appName"`
	/**
	*终端平台 [ANDROID,IOS]
	**/
	Platform string `json:"platform"`
	/**
	*android包名 (android必填)
	**/
	PackageID string `json:"packageID"`
	/**
	*设备ID
	**/
	DeviceID string `json:"deviceID"`
	/**
	*授权SDK名称: [haixin,esandinfo]
	**/
	SdkName string `json:"sdkName"`
	/**
	*授权SDK版本
	**/
	SdkVersion string `json:"sdkVersion"`
}
/**
*一砂云请求action
**/
const GetCertificateRequestAct = "comms/efaceid/getCertificate"
/**
 * 转换成JSON字符串
 */
 func (self GetCertificateRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}