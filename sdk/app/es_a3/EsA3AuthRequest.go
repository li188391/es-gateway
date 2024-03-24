package es_a3

import(
	"encoding/json"
)

type EsA3AuthRequest struct{
	/**
	*公共参数bizId相同即可
	**/
	TransId string `json:"transId"`
	/**
	*是否需要对图片进行活体检测(主要针对照片翻拍场景，默认为false)
	**/
	HadLiving bool `json:"hadLiving"`
	/**
	*身份证号码
	**/
	IdNumber string `json:"idNumber"`
	/**
	*姓名
	**/
	Name string `json:"name"`
	/**
	*人像照片数据（Base64URLSafe编码字符串）注：人像照片的大小不能大于40kb
	**/
	PhotoData string `json:"photoData"`
}
/**
*一砂云，请求action
**/
const EsA3AuthRequestAct = "ctid/es_a3/auth"
/**
 * 转换成JSON字符串
 */
 func (self EsA3AuthRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
