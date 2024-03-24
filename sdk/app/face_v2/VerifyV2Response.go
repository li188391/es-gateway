package face_v2

import(
	"encoding/json"
)

type VerifyV2Response struct{
	/**
	*请求id，用于关联请求日志
	**/
	RequestId string `json:"requestId"`
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*响应描述
	**/
	Msg string `json:"msg"`
	/**
	*置信值，Float类型，取值［0，100］，数字越大表示风险越小。建议 75 分通过
	**/
	Confidence float64 `json:"confidence"`
}
/**
 * 转换成JSON字符串
 */
 func (self VerifyV2Response) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}