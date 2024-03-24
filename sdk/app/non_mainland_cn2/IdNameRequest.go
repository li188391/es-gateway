package non_mainland_cn2

import(
	"encoding/json"
)

type IdNameRequest struct{
	/**
	*国籍，华侨和港澳人员用 CHN (默认值)
	**/
	Nation string `json:"nation"`
	/**
	*证件ID类型，当前支持：
    *港澳居民来往大陆通行证/回乡证 （3）
    *台湾居民来往大陆通行证/回乡证 （4）
    *中国大陆护照 （5）
    *中国大陆永久居住证  （6）
	**/
	IdType string `json:"idType"`
	/**
	*姓名
	**/
	Name string `json:"name"`
	/**
	*证件号码
	**/
	IdNo string `json:"idNo"`
}
/**
*一砂云请求action
**/
const IdNameRequestAct = "comms/identity_auth/non_mainland_cn/id_name"
/**
 * 转换成JSON字符串
 */
 func (self IdNameRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
