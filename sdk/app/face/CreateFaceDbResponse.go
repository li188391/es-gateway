package face

import(
	"encoding/json"
)

type CreateFaceDbResponse struct{
	/**
     * 状态码
    **/
	Code string `json:"code"`
	/**
	*请求id
	**/
	RequestId string `json:"requestId"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	
	
}
/**
 * 转换成JSON字符串
 */
 func (self CreateFaceDbResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}