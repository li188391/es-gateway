package ifaa

import(
	"encoding/json"
)

type SAuthRequest struct{
	/**
	*ifaa认证数据 (由app生成)
	**/
	IfaaMsg string `json:"ifaaMsg"`

} 
/**
*阿里云请求URL
**/
const SAuthRequestALIYUN_URL = "http://simpleifaa.market.alicloudapi.com/ifaa/sAuth"
/**
 * 转换成JSON字符串
 */
 func (self SAuthRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
