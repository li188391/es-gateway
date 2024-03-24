package rpc

import(
	"encoding/json"
)

type GetBehaviorRequest struct{
	
} 
/**
*阿里云请求URL
**/
const GetBehaviorRequestALIYUN_URL = "https://edis3v.market.alicloudapi.com/getBehavior"
/**
 * 转换成JSON字符串
 */
 func (self GetBehaviorRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
