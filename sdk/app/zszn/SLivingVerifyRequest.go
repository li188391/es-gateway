package zszn

import(
	"encoding/json"
)

type SLivingVerifyRequest struct{
	/**
	*被检测照片数据（Base64编码后），人脸图像最长边像素为800pi最佳
	**/
	Image string `json:"image"`
	
}
/**
*阿里云请求URL
**/
const SLivingVerifyRequestALIYUN_URL = "http://zsverify.market.alicloudapi.com/comms/zszn/sLivingVerify"
/**
 * 转换成JSON字符串
 */
 func (self SLivingVerifyRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}