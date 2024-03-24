package smsall

import(
	"encoding/json"
)

type SendmsgAllRequest struct{
	/**
	*电话号码，号码前必须加国家区号，如中国号码需在前面加86，印度加91，美国加1
	**/
	Mobile string `json:"mobile"`
	/**
	*模板 ID, 登录 https://openali.esandcloud.com申请
	**/
	TemplateID string `json:"templateID"`
	/**
	*模板参数，多个参数以逗号“,”分隔，入股对应的模板无参数此字段可为空
	**/
	TemplateParamSet string `json:"templateParamSet"`
	/**
	*短信下发状态回调通知地址（POST主动回调,需要在业务服务器留接口）
	**/
	CallbackUrl string `json:"callbackUrl"`
}
/**
*一砂云，请求action
**/
const SendmsgAllRequestAct = "comms/sms/sendmsgall"
/**
 * 转换成JSON字符串
 */
 func (self SendmsgAllRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}