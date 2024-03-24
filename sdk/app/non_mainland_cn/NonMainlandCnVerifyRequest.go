package non_mainland_cn

import(
	"encoding/json"
)

type NonMainlandCnVerifyRequest struct{
	/**
	*认证token
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
const NonMainlandCnVerifyRequestAct = "livingdetection/non_mainland_cn/verify"
/**
 * 转换成JSON字符串
 */
 func (self NonMainlandCnVerifyRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
