package zfi

import(
	"encoding/json"
)

type GetResultResponse struct{
	/**
	*业务号（与初始化传过来的一致）
	**/
	BizId string `json:"bizId"`
	/**
	*业务ID
	**/
	BusinessId string `json:"businessId"`
	/**
	*请求ID(用于查日志)
	**/
	RequestId string `json:"requestId"`
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*人像置信度（区间：0-1000）阈值参考如下：
	*错误接受率1%，阈值比对分数500 
    *错误接受率0.1%，阈值比对分数600
	*错误接受率0.01%，阈值比对分数700
	*错误接受率0.001%，阈值比对分数800
	*错误接受率0.0001%，阈值比对分数900
	*推荐阈值比对分数700以上
	**/
	Confidence string `json:"confidence"`
	/**
	*认证结果：T：实人成功；F：实人失败
	**/
	Passed string `json:"passed"`
	/**
	*实人结果码：
	**/
	ResultCode string `json:"resultCode"`
	/**
	*结果描述
	**/
	ResultMsg string `json:"resultMsg"`
	/**
	*活体类型：0: 混合版 1：静默活体；2：读数活体 
	**/
	Type string `json:"type"`
	/**
	*认证用户姓名
	**/
	CertName string `json:"certName"`
	/**
	*认证用户身份证号
	**/
	CertNo string `json:"certNo"`
	/**
	*活体组合（静默活体时生效） 1：远近，3：摇头，4: 点头
	**/
	LivingType string `json:"livingType"`
	/**
	*活体检测视频地址(有效时间为30分钟)
	**/
	VideoUrl string `json:"videoUrl"`
	/**
	*活体检测最清晰的一张照片，有效期，30分钟
	**/
	ImgUrl string `json:"imgUrl"`
}
/**
 * 转换成JSON字符串
 */
 func (self GetResultResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}