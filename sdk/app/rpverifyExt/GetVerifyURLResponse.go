package rpverifyExt

import(
	"encoding/json"
)

type GetVerifyURLResponse struct{
	/**
	*错误码
	**/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*认证链接
	**/
	Url string `json:"url"`
}
/**
 * 转换成JSON字符串
 */
 func (self GetVerifyURLResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}