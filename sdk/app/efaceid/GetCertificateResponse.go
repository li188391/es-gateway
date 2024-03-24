package efaceid

import(
	"encoding/json"
)

type GetCertificateResponse struct{
	/**
	*错误码
	**/
	Code string `json:"code"`
	/**
	*新增成功的返回信息
	**/
	Msg string `json:"msg"`
	/**
	*证书数据 (base64 编码)
	**/
	Certificate string `json:"certificate"`
	/**
	*起始时间
	**/
	NotBefore string `json:"notBefore"`
	/**
	*截止时间
	**/
	NotAfter string `json:"notAfter"`
}
/**
 * 转换成JSON字符串
 */
 func (self GetCertificateResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}