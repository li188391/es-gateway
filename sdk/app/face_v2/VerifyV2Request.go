package face_v2

import(
	"encoding/json"
)

type VerifyV2Request struct{
	/**
	*客户业务id(32位uuid)，保证其唯一
	**/
	BizId string `json:"bizId"`
	/**
	*人脸对比源图片base64字符串
	**/
	SourceImage string `json:"sourceImage"`
	/**
	*待验证的人脸图片base64字符串
	**/
	VerifyImage string `json:"verifyImage"`
}
/**
*一砂云，请求action
**/
const VerifyV2RequestAct = "faceId/face/verify_v2"
/**
*阿里云网关URL
**/
const VerifyV2RequestALIYUN_URL = "http://efacecmp.market.alicloudapi.com/face_verify"
/**
 * 转换成JSON字符串
 */
 func (self VerifyV2Request) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}