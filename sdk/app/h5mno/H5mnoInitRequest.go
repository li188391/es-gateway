package h5mno

import(
	"encoding/json"
)

type H5mnoInitRequest struct{
	/**
	*业务ID（应保证其唯一）
	**/
	BizId string `json:"bizId"`
	/**
    * 手机号
    */
	PhoneNumber string `json:"phoneNumber"`
	/**
	*回调通知url(认证完成后回调此接口)
	**/
	ReturnUrl string `json:"returnUrl"`
	
}
/**
*阿里云请求URL
**/
const H5mnoInitRequestALIYUN_URL = "http://h5mno.market.alicloudapi.com/init"
/**
 * 转换成JSON字符串
 */
 func (self H5mnoInitRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}