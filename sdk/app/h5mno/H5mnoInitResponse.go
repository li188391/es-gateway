package h5mno

import(
	"encoding/json"
)

type H5mnoInitResponse struct{
	/**
    * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*认证url
	**/
	Url string `json:"url"`
}
/**
 * 转换成JSON字符串
 */
 func (self H5mnoInitResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}