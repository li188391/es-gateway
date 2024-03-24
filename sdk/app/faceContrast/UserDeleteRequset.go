package faceContrast

import(
	"encoding/json"
)

type UserDeleteRequest struct{
	/**
	*用户唯一ID
	**/
	Uuid string `json:"uuid"`
}
/**
*一砂云，请求action
**/
const UserDeleteRequestAct = "livingdetection/faceContrast/user/delete"
/**
 * 转换成JSON字符串
 */
 func (self UserDeleteRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
