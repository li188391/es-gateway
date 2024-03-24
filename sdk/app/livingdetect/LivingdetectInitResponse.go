package livingdetect

import(
	"encoding/json"
)

type LivingfdetectResponse struct{
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*TOKEN
	**/
	Token string `json:"token"`
}
/**
 * 转换成JSON字符串
 */
 func (self LivingfdetectResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}