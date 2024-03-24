package ocr_SDK

import(
	"encoding/json"
)

type OcrInitRequest struct{
	/**
	*ocr类型：
     0: 身份证正面，
     1: 身份证反面，
	**/
	OcrType string  `json:"ocrType"`
}
/**
*一砂云，请求action
**/
const OcrInitRequestAct = "ocr/init"
/**
 * 转换成JSON字符串
 */
 func (self OcrInitRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}