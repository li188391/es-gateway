package edis_ctid2

import(
	"encoding/json"
)

type IdNameRequest struct{
	/**
	*姓名
	**/
	Name string `json:"name"`
	/**
	*身份证号码
	**/
	IdCard string `json:"idCard"`
	/**
	*证件类型,CN: 中国大陆身份证(默认值)，CN_PT: 中国大陆护照，CN_RT: 中国大陆永久居住证，TW: 中国台湾通行证，HK: 中国香港通行证，MO: 中国澳门通行证
	**/
	IdcardType string `json:"IdcardType"`
	
}  
/**
*阿里云请求URL
**/
const IdNameRequestALIYUN_URL = "http://edis2s.market.alicloudapi.com/edis_ctid_id_name"
/**
 * 转换成JSON字符串
 */
 func (self IdNameRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
