package rpverify

import(
	"encoding/json"
)

type RpverifyVerifyRequest struct{
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
*一砂云请求action
**/
const RpverifyVerifyRequestAct = "livingdetection/rpverify/verify"
/**
 * 转换成JSON字符串
 */
 func (self RpverifyVerifyRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
