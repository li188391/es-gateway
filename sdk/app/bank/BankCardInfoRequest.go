package bank

import(
	"encoding/json"
)

type BankCardInfoRequest struct{
	/**
	*银行卡号
	**/
	AccountNo string `json:"accountNo"`
}  
/**
*阿里云请求URL
**/
const BankCardInfoRequestALIYUN_URL = "http://ebankcard.market.alicloudapi.com/bankCardInfo"
/**
 * 转换成JSON字符串
 */
 func (self BankCardInfoRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
