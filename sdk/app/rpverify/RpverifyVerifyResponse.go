package rpverify

import(
	"encoding/json"
)

type RpverifyVerifyResponse struct{
	/**
	*执行结果,参考状态码
	**/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*业务ID(和初始化传入的一致)
	**/
	BizId string `json:"bizId"`
	/**
	*请求ID (用于查日志)
	**/
	RequestId string `json:"requestId"`
	/**
	*活体检测类型：1：远近，2：眨眼，3：摇头，4: 点头，5:张嘴，6：炫彩，7: 读数
	**/
	LivingType string `json:"livingType"`
	/**
	*姓名
	**/
	CertName string `json:"certName"`
	/**
	*身份证号
	**/
	CertNo string `json:"CertNo"`
	/**
	*活体认证截取的最清晰的一张用户照片的下载url地址，有效期为 24 个小时
	**/
	BestImg string `json:"bestImg"`
	/**
	*实人认证结果,T： 实人认证成功（分数>=750分）F：实人认证失败（分数<750分）
	**/
	Pass string `json:"pass"`
	/**
	*实人认证的置信值，分数在1-1000,建议750分算通过
	**/
	Rxfs string `json:"rxfs"`
}
/**
 * 转换成JSON字符串
 */
 func (self RpverifyVerifyResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}