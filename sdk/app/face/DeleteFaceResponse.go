package face

import(
	"encoding/json"
)

type DeleteFaceResponse struct{
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*请求id
	**/
	RequestId string `json:"requestId"`
	/**
	*响应描述
	**/
	Msg string `json:"msg"`
}
/**
 * 转换成JSON字符串
 */
 func (self DeleteFaceResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}