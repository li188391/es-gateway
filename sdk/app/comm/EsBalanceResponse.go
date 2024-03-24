package comm

import "encoding/json"

type EsBalanceResult struct {
	/**
	* 一砂云账号ID
	 */
	UserId string `json:"userId"`
	/**
	* 账户余额
	 */
	Balance string `json:"balance"`
	/**
	* 最大授信额度
	 */
	MaximumCredit float64 `json:"maximumCredit"`
	/**
	* 剩余授权额度
	 */
	RemainingAuthorizationAmount float64 `json:"remainingAuthorizationAmount"`
	/**
	* 欠费阈值 (达到这个阈值将进行短信/邮件通知)
	 */
	Threshold float64 `json:"threshold"`
}

/**
 * 查询余额响应对象
 */
type EsBalanceResponse struct {
	/**
	 * 日志ID
	 */
	LogId string `json:"logId"`
	/**
	 * 响应状态码
	 */
	Code string `json:"code"`
	/**
	 * 执行结果描述
	 */
	Msg string `json:"msg"`
	/**
	 * 查询结果 (包括子账号的)
	 */
	Result []EsBalanceResult `json:"result"`
}

/**
 * 转换成JSON字符串
 */
func (self EsBalanceResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
