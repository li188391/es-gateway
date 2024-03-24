package faceId

import(
	"encoding/json"
)

type GetVerifyTokenRequest struct{
	/**
	*客户业务流水号，该号必须唯一。并会在notify和return时原封不动的返回给您的服务器，以帮助您确认每一笔
	**/
	BizNo string `json:"bizNo"`
	/**
	*认证姓名，和idNumber必须同时为空或不为空，如不为空，那么将跳过用户身份录入流程
	**/
	IdName string `json:"idName"`
	/**
	*认证身份证号码，和idName必须同时为空或不为空，如不为空，那么将跳过用户身份录入流程
	**/
	IdNumber string `json:"idNumber"`
	/**
	*用户完成或取消验证后网页跳转的目标URL。（回调方法为Post）
	**/
	NotifyUrl string `json:"notifyUrl"`
	/**
	*OCR是否包括身份证背面, 默认为false
	**/
	OcrIncIdBack string `json:"ocrIncIdBack"`
	/**
	*是否值支持OCR, 默认为false
	**/
	OcrOnly string `json:"ocrOnly"`
	/**
	*页面背景颜色
	**/
	PageBgColor string `json:"pageBgColor"`
	/**
	*是否返回ocr上传的身份证照片, 默认为false
	**/
	RetIdImg string `json:"retIdImg"`
	/**
	*是否返回用户照片，默认为false
	**/
	ReturnImg string `json:"returnImg"`
	/**
	*是否返回视频活体下载url以及认证照片下载url，默认为false
	**/
	ReturnLiveVideo string `json:"returnLiveVideo"`
	/**
	*用户完成验证、取消验证、或验证超时后，由FaceID服务器请求客户服务器的URL。（推荐为HTTPS协议，如果为HTTP则用户需要通过签名自行校验数据可信性，回调方法为Post）
	**/
	ReturnUrl string `json:"returnUrl"`
	/**
	*页面按钮背景颜色
	**/
	TxtBgColor string `json:"txtBgColor"`

}
/**
*一砂云，请求action
**/
const GetVerifyTokenRequestAct = "faceId/get_verify_token"
/**
 * 转换成JSON字符串
 */
 func (self GetVerifyTokenRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}