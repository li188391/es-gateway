package comm

import (
	"encoding/json"
)

type EsBalanceRequest struct {
	/**
	 * 请求业务ID（需确保唯一）
	 */
	TransId string `json:"transId"`
}

/**
 * 一砂云对应的ACTION值
 */
const EsBalanceRequestAct = "comms/edis/query_balance_info"

func (self EsBalanceRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
