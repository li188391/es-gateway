package edis_ctid3

import(
	"encoding/json"
)

type IdNameImageRequest struct{
	/**
	*身份证号码
	**/
	IdCard string `json:"idCard"`
	/**
	*姓名
	**/
	Name string `json:"name"`
	/**
	*人像数据（Base64URLSafe编码字符串，注：人像照片的大小不能大于40kb）
	**/
	Image string `json:"image"`
	/**
	*是否需要对图片进行活体检测(主要针对照片翻拍场景，默认为false)
	**/
	HadLiving string `json:"hadLiving"`
	/**
	*活体严格级别, loose:宽松,standard:标准,strict:严格
	**/
	SecLevel string `json:"secLevel"`
}  
/**
*阿里云请求URL
**/
const IdNameImageRequestALIYUN_URL = "http://edis3p.market.alicloudapi.com/edis_ctid_id_name_image"
/**
 * 转换成JSON字符串
 */
 func (self IdNameImageRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
