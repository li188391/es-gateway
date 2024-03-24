package ocr_SDK

import(
	"encoding/json"
)

type OcrInitResponse struct{
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*此次请求ID
	**/
	RequestId string `json:"requestId"`
	/**
	*请求TOKEN
	**/
	Token string `json:"token"`
}
/**
 * 转换成JSON字符串
 */
 func (self OcrInitResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}