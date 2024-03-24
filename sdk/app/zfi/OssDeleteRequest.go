package zfi

import(
	"encoding/json"
)

type OssDeleteRequest struct{
	/**
	*初始化返回的业务id（注意切勿让中间人替换此ID）
	**/
	BusinessId string `json:"businessId"`
}
/**
*一砂云，请求action
**/
const OssDeleteRequestAct = "comms/zfi/oss/delete"
/**
*阿里云网关URL
**/
const OssDeleteRequestALIYUN_URL = "https://zimfaceid1.market.alicloudapi.com/comms/zfi/rmCachedData"
/**
 * 转换成JSON字符串
 */
 func (self OssDeleteRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}