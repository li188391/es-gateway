package edis_ctid3


import(
	"encoding/json"
)

type IdNameImageResponse struct{
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*业务ID,用于关联业务及日志
	**/
	TransId string `json:"transId"`
	/**
	*人脸比对可信度，总分1000,默认750通过
	**/
	Rxfs string `json:"rxfs"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	
}
/**
 * 转换成JSON字符串
 */
 func (self IdNameImageResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}