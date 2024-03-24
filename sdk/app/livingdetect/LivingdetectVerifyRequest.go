package livingdetect

import(
	"encoding/json"
)

type LivingdetectVerifyRequest struct{
	/**
	*token(初始化时候返回)
	**/
	Token string `json:"token"`
	/**
	*认证数据（活体检测完成后从sdk返回）
	**/
	VerifyMsg string `json:"verifyMsg"`
}
/**
*一砂云请求action
**/
const LivingdetectVerifyRequestAct = "livingdetection/livingdetect/verify"
/**
 * 转换成JSON字符串
 */
 func (self LivingdetectVerifyRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}