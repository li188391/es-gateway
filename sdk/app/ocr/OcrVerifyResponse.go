package ocr

import (
	"encoding/json"
)

type OcrVerifyResponse struct{
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
	*OCR响应数据
	**/
	Data float64 `json:"data"`
}
/**
 * 转换成JSON字符串
 */
 func (self OcrVerifyResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}