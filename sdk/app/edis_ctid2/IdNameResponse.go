package edis_ctid2


import(
	"encoding/json"
)

type IdNameResponse struct{
	/**
	*版本号
	**/
	Ver string `json:"ver"`
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*生日
	**/
	Birthday string `json:"birthday"`
	/**
	*性别
	**/
	Sex string `json:"sex"`
	/**
	*业务ID
	**/
	TransId string `json:"transId"`	
}
/**
 * 转换成JSON字符串
 */
 func (self IdNameResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}