package livingdetect

import(
	"encoding/json"
)

type LivingdetectRequest struct{
	/**
	*初始化信息 (从SDK端生成)
	**/
	InitMsg string `json:"initMsg"`

} 
/**
*一砂云请求action
**/
const LivingdetectRequestAct = "livingdetection/livingdetect/init"
/**
 * 转换成JSON字符串
 */
 func (self LivingdetectRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
