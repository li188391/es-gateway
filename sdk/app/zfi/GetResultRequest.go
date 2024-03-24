package zfi

import(
	"encoding/json"
)

type GetResultRequest struct{
	/**
	*初始化返回的业务id（注意切勿让中间人替换此ID）
	**/
	BusinessId string `json:"businessId"`
}
/**
*一砂云，请求action
**/
const GetResultRequestAct = "comms/zfi/get_result"
/**
*阿里云网关URL
**/
const GetResultRequestALIYUN_URL = "https://zimfaceid1.market.alicloudapi.com/comms/zfi/get_result"
/**
 * 转换成JSON字符串
 */
 func (self GetResultRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}