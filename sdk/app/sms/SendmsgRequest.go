package sms

import(
	"encoding/json"
)

type SendmsgRequest struct{
	/**
	*中国大陆电话号码
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
	*短信下发状态回调通知地址（POST主动回调）
	**/
	CallbackUrl string `json:"callbackUrl"`
}
/**
*一砂云，请求action
**/
const SendmsgRequestAct = "comms/sms/sendmsg"
/**
 * 转换成JSON字符串
 */
 func (self SendmsgRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}