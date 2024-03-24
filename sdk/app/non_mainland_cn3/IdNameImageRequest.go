package non_mainland_cn3

import(
	"encoding/json"
)

type IdNameImageRequest struct{
	/**
	*国籍，华侨和港澳人员用 CHN (默认值)
	**/
	Nation string `json:"nation"`
	/**
	*证件ID类型，当前支持：
    *1. 港澳通行证/回乡证
    *5. 中国大陆护照
	**/
	IdType string `json:"idType"`
	/**
	*姓名
	**/
	Name string `json:"name"`
	/**
	*证件号码
	**/
	IdNO string `json:"idNO"`
	/**
	*人像照片BASE64字符串
	**/
	PersonImg string `json:"personImg"`
}
/**
*一砂云请求action
**/
const IdNameImageRequestAct = "comms/identity_auth/non_mainland_cn/id_name_image"
/**
 * 转换成JSON字符串
 */
 func (self IdNameImageRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}