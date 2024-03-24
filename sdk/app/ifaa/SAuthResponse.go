package ifaa

import(
	"encoding/json"
)

type SAuthResponse struct{
	/**
    * IFAA协议版本，当前为 2.0.0
    **/
	Ver string `json:"ver"`
	/**
    * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*生物认证认证类型，可参考: 备注.认证类型
	**/
	AuthType string `json:"authType"`
	/**
	*TOKEN, 全局唯一
	**/
	Token string `json:"token"`
	/**
	*指纹ID(某个指位的唯一ID)/指纹集合ID, 对于IOS和部分Android手机是指纹集合ID, 有新录入指纹此id会变更。
	**/
	Id string `json:"id"`
	/**
	*生物特征信息
	**/
	BioInfo string `json:"bioInfo"`
}
/**
 * 转换成JSON字符串
 */
 func (self SAuthResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}