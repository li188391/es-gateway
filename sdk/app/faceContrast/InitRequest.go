package faceContrast

import(
	"encoding/json"
)

type InitRequest struct{
	/**
	*业务ID(必须保证唯一，方便后续业务跟踪)
	**/
	BizId string `json:"bizId"`
	/**
	*初始化信息(从SDK端生成)
	**/
	InitMsg string `json:"initMsg"`
}
/**
*一砂云,请求action
**/
const InitRequestAct = "livingdetection/faceContrast/init"
/**
 * 转换成JSON字符串
 */
 func (self InitRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}