package zszn

import(
	"encoding/json"
)

type MobileStatusRequest struct{
	/**
	*手机号，可传多个以逗号为分隔符(最多500个),如123,456，如果多个号码，结果将通过notifyUrl进行异步回调
	**/
	Mobile string `json:"mobile"`
	/**
    * 批量查询时回调URL（POST请求，报文格式参考返回示例）
    */
	NotifyUrl string `json:"notifyUrl"`
	
}
/**
*阿里云请求URL
**/
const MobileStatusRequestALIYUN_URL = "http://edisonline.market.alicloudapi.com/comms/zszn/mobileStatus"
/**
 * 转换成JSON字符串
 */
 func (self MobileStatusRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}