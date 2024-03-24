package rpverify

import(
	"encoding/json"
)

type RpverifyInitRequest struct{
	/**
	*业务ID(必须保证唯一，方便后续业务跟踪)
	**/
	BizId string `json:"bizId"`
	/**
	*姓名
	**/
	CertName string `json:"certName"`
	/**
	*身份证号
	**/
	CertNo string `json:"certNo"`
	/**
	*初始化信息(从SDK端生成)
	**/
	InitMsg string `json:"initMsg"`
}
/**
*一砂云请求action
**/
const RpverifyInitRequestAct = "livingdetection/rpverify/init"
/**
 * 转换成JSON字符串
 */
 func (self RpverifyInitRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
