package face

import(
	"encoding/json"
)

type EsBalanceResult struct{
	/**
	*其中一个匹配人脸数据对象
	**/
	FaceItems struct{
		/**
		*实体ID
		**/
		EntityId string `json:"entityId"`
		/**
		人脸的相似度，取值范围-1.0~1.0，取小数点后6位，数值越大相似度越高。
		**/
		Score float64 `json:"score"`
		/**
		*自定义信息
		**/
		ExtraData string `json:"extraData"`

	} `json:"faceItems"`
	/**
	*人脸左上角y值
	**/
	Location  struct {
		/**
		*人脸宽度
		**/
		Width  float64 `json:"width"`
		/**
		*人脸高度
		**/
		Height float64 `json:"height"`
		/**
		*人脸左上角x值
		**/
		X float64 `json:"x"`
		/**
		人脸左上角y值
		**/
		Y float64 `json:"y"`
	} `json:"location"`

}

type SearchFaceResponse struct{
	/**
     * 响应状态码
    **/
	Code string `json:"code"`
	/**
	*请求id
	**/
	RequestId string `json:"requestId"`
	/**
	*响应描述
	**/
	Msg string `json:"msg"`
	/**
	*匹配数据数组
	**/
	MatchList []EsBalanceResult `json:"matchList"`
	
}
/**
 * 转换成JSON字符串
 */
 func (self SearchFaceResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}