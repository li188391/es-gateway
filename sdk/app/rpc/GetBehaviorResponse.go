package rpc

import(
	"encoding/json"
)

type GetBehaviorResponse struct{
	/**
    * 版本号
    **/
	Ver string `json:"ver"`
	/**
	*业务Id
	**/
	TransId string `json:"transId"`
	/**
    * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*活体行为
	**/
	Behavior string `json:"behavior"`
	/**
	*活体行为类型
	**/
	BehaviorType string `json:"behaviorType"`
	/**
	*返回token
	**/
	BehaviorToken string `json:"behaviorToken"`
	
}
/**
 * 转换成JSON字符串
 */
 func (self GetBehaviorResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}