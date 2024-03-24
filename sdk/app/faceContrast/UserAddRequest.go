package faceContrast

import(
	"encoding/json"
)

type UserAddRequest struct{
	/**
	*用户唯一ID
	**/
	Uuid string `json:"uuid"`
	/**
	*图片BASE64字符串，与verifyMsg, imageURL三选其一
	**/
	Image string `json:"image"`
	/**
	*图片URL地址，与image, verifyMsg三选其一
	**/
	ImageURL string `json:"imageURL"`
	/**
	*认证TOKEN （选择verify时传入)
	**/
	Token string `json:"token"`
	/**
	*活体检测数据，与image, imageURL三选其一
	**/
	VerifyMsg string `json:"VerifyMsg"`
}
/**
*一砂云请求action
**/
const UserAddRequestAct = "livingdetection/faceContrast/user/add"
/**
 * 转换成JSON字符串
 */
 func (self UserAddRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}