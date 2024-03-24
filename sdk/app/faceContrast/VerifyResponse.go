package faceContrast

import(
	"encoding/json"
)

type VerifyResponse struct{
	/**
	*错误码
	**/
	Code string `json:"code"`
	/**
	*活体检测结果
	**/
	IdtResult string `json:"ldtResult"`
	/**
	*人脸比对结果
	**/
	FaceVerifyResult string `json:"faceVerifyResult"`
	/**
	*结果描述
	**/
    Msg string `json:"msg"`
	/**
	*人脸比对是否通过
	**/
	Pass string `json:"passed"`
	/**
	*人脸比对分数（通常认为大于75分为同一个人）
	**/
	Confident string `json:"confident"`
	/**
	*请求ID
	**/
	RequestId string `json:"requestId"`
	/**
	*认证token
	**/
	Token string `json:"token"`
	/**
	*活体检测的截图
	**/
	BestImg string `json:"bestImg"`
	/**
	*活体检测类型
	**/
	LivingType string `json:"livingType"`
}
/**
 * 转换成JSON字符串
 */
 func (self VerifyResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}