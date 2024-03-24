package non_mainland_cnH5
import(
	"encoding/json"
)

type GetAuthUrlRequest struct{
	/**
	*业务id，务必保证唯一UUID，建议传入
	**/
	BizId string `json:"bizId"`
	/**
	*静默活体时生效，可组合活体动作类型，如13为先远近后摇头 :
     1：远近
     3：摇头
     4：点头
     6：炫彩
     7：读数
	 **/
	LivingType string `json:"livingType"`
	/**
	*证件类型， 
     1. 港澳通行证/回乡证 
     5. 中国大陆护照
	**/
	IdType string `json:"idType"`
	/**
	*姓名
	**/
	CertName string `json:"certName"`
	/**
	*证件ID(港澳通行证ID，护照ID)
	**/
	CertNo string `json:"certNo"`
	/**
	*回调重定向地址（认证完成后，从浏览器端重定向到此页面，并把认证结果以表单的形式提交到服务器，建议对结果进行验签）
	**/
	ReturnUrl string `json:"returnUrl"`
	/**
	*认证完成后，结果回调地址(一砂服务器调用业务服务器，建议对结果进行验签)
	**/
	NotifyUrl string `json:"notifyUrl"`
}
/**
*一砂云请求action
**/
const GetAuthUrlRequestAct = "livingdetection/nonMainland/getAuthUrl"
/**
 * 转换成JSON字符串
 */
 func (self GetAuthUrlRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
