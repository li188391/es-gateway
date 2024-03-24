package rpverifyExt

import(
	"encoding/json"
)

type GetVerifyURLRequest struct{
	/**
	*对应的证件ID
	**/
	Eid string `json:"eid"`
	/**
	*执行完成后，跳转页面 (在query中附带执行结果 )
	*/
	ReturnUrl string `json:"returnUrl"`
	/**
	*执行完成后，结果回调地址
	**/
	NotifyUrl string `json:"notifyUrl"`
	/**
	*活体检测组合, 如：13 为先远近检测，后摇头检测: 
    *1：远近
    *3：摇头
    *4:  点头
    *6：炫彩
	**/
	LivingType string `json:"livingType"`
}
/**
*一砂云请求action
**/
const GetVerifyURLRequestAct = "livingdetection/rpverifyExt/getVerifyURL"
/**
 * 转换成JSON字符串
 */
 func (self GetVerifyURLRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}