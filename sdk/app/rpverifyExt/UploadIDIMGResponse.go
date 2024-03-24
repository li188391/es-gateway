package rpverifyExt

import (
	"encoding/json"
)

type UploadIDIMGResponse struct{
	/**
	*错误码
	**/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*证件ID（唯一）
	**/
	Eid string `json:"eid"`
}
/**
 * 转换成JSON字符串
 */
 func (self UploadIDIMGResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}