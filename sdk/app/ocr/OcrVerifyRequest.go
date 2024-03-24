package ocr

import (
	"encoding/json"
)

type OcrVerifyRequest struct{
	/**
	*照片数据（Base64编码后）
	**/
	File string `json:"file"`
	/**
	*ocr类型：
     0: 身份证正面，
     1: 身份证反面，
     2: 银行卡
     3：营业执照
	**/
	OcrType string `json:"ocrType"`
}
/**
*一砂云，请求action
**/
const OcrVerifyRequestAct = "ocr/verify"
/**
* 阿里云请求URL
**/
const OcrVerifyRequestALIYUN_URL = "http://edisocr1.market.alicloudapi.com/ocr/verify"
/**
 * 转换成JSON字符串
 */
 func (self OcrVerifyRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}