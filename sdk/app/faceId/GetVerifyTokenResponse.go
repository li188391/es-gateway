package faceId

import(
	"encoding/json"
)

type GetVerifyTokenResponse struct{
	/**
	*业务号
	**/
	BizNo string `json:"bizNo"`
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*超时时间(时间戳，有效期为30分钟)
	**/
	Expired string `json:"expired"`
	/**
	*认证链接
	**/
	VerifyUrl string `json:"verifyUrl"`
}
/**
 * 转换成JSON字符串
 */
 func (self GetVerifyTokenResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}