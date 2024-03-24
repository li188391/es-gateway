package zszn

import(
	"encoding/json"
)

type MobileOnlineIntervalRequest struct{
	/**
	*手机号
	**/
	Mobile string `json:"mobile"`
	
}
/**
*阿里云请求URL
**/
const MobileOnlineIntervalRequestALIYUN_URL = "http://edistime.market.alicloudapi.com/comms/zszn/mobileOnlineInterval"
/**
 * 转换成JSON字符串
 */
 func (self MobileOnlineIntervalRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}