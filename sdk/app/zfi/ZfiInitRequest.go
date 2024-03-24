package zfi

import(
	"encoding/json"
)

type ZfiInitRequest struct{
	/**
	*业务id，务必保证唯一
	**/
	BizId string `json:"bizId"`
	/**
	*姓名
	**/
	IdName string `json:"idName"`
	/**
	*身份证号
	**/
	IdNumber string `json:"idNumber"`
	/**
	*回调重定向地址（浏览器前端以表单的形式提交（post））
	**/
	ReturnUrl string `json:"returnUrl"`
	/**
	*异步通知地址（一砂服务器直接以post的方式推送）
	**/
	NotifyUrl string `json:"notifyUrl"`
	/**
	*活体策略类型 （建议留空，或者填入 0 ）：
    0:（默认值）优先选择交互活体，如无法拉起摄像头选择读数活体
    1：只用交互活体 (忽略浏览器兼容性问题)；
    2：只用读数活体
	**/
	Type string `json:"type"`
	/**
	*活体检测组合，如：13 为先远近检测，后摇头检测 : 
    1：远近，
    3：摇头，
    4: 点头，
    6：炫彩
	**/
	LivingType string `json:"livingType"`
	/**
	*是否需要返回活体检测视频，true:返回活体视频，false:不返回活体视频，默认false
	**/
	NeedVideo string `json:"needVideo"`
}
/**
*一砂云，请求action
**/
const ZfiInitRequestAct = " comms/zfi/init"
/**
*阿里云，请求URL
**/
const ZfiInitRequestALIYUN_URL = "https://zimfaceid1.market.alicloudapi.com/comms/zfi/init"
/**
 * 转换成JSON字符串
 */
 func (self ZfiInitRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}