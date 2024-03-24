package rpc

import(
	"encoding/json"
)

type VerifyResponse struct{
	/**
    * 版本号
    **/
	Ver string `json:"ver"`
	/**
	*业务Id
	**/
	TransId string `json:"transId"`
	/**
    * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*图片信息
	**/
	ImageBest string `json:"imageBest"`
	/**
	*人像比对分数
	**/
	Rxfs string `json:"rxfs"`	
	
}
/**
 * 转换成JSON字符串
 */
 func (self VerifyResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}