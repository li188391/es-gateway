package faceContrast

import(
	"encoding/json"
)

type VerifyRequest struct{
	/**
	*比对照片数据(base64字符串)
	**/
	Image string `json:"image"`
	/**
	*用户注册ID 
	**/
	Uuid string `json:"uuid"`
	/**
	*比对照片的URL地址 (确保地址公网可访问)
	**/
	ImageUrl string `json:"imageUrl"`
	/**
	*认证token(初始化时候返回)
	**/
	Token string `json:"token"`
	/**
	*认证数据（认证完成后从sdk返回）
	**/
	VerifyMsg string `json:"verifyMsg"`
}
/**
*一砂云，请求action
**/
const VerifyRequestAct = "livingdetection/faceContrast/verify"
/**
 * 转换成JSON字符串
 */
 func (self VerifyRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}