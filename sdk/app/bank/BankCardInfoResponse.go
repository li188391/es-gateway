package bank


import(
	"encoding/json"
)

type BankCardInfoResponse struct{
	/**
	*请求ID，用于关联请求日志
	**/
	RequestId string `json:"requestId"`
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*格式是否正确
	**/
	Validated bool `json:"validated"`
	/**
	 * 银行卡类型
     * DC: '储蓄卡',
     * CC: '信用卡',
     * SCC: '准贷记卡',
     * PC: '预付费卡'
	**/
	CardType string `json:"cardType"`
	/**
	*银行名称
	**/
	BankName string `json:"bankName"`	
	/**
	*银行代码
	**/
	BankCode string `json:"bankCode"`	
	/**
	*是否收费：true 收费，false 未收费
	**/
	Charge string `json:"charge"`	
}
/**
 * 转换成JSON字符串
 */
 func (self BankCardInfoResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}