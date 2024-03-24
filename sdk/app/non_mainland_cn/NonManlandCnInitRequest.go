package non_mainland_cn

import(
	"encoding/json"
)

type NonMainlandCnRequest struct{
	/**
	*业务ID(必须保证唯一，方便后续业务跟踪)
	**/
	BizId string `json:"bizId"`
	/**
	*初始化信息(从SDK端生成)
	**/
	InitMsg string `json:"initMsg"`
	/**
	*国籍、地区编码，华侨和港澳人员使用CHN
	**/
	Nation string `json:"nation"`
	/**
	*证件类型，1. 港澳通行证/回乡证；5. 中国大陆护照
	**/
    IdType string `json:"idType"`
	/**
	*姓名
	**/
	CertName string `json:"certName"`
	/**
	*证件号码
	**/
	CertNo string `json:"certNo"`
}
/**
*一砂云请求action
**/
const NonMainlandCnRequestAct = "livingdetection/non_mainland_cn/init"
/**
 * 转换成JSON字符串
 */
 func (self NonMainlandCnRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
