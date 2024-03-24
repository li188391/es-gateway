package non_mainland_cn

import(
	"encoding/json"
)

type NonMainlandCnVerifyResponse struct{
	/**
	*请求ID
	**/
	RequestId string `json:"requestId"`
	/**
	*业务ID(和初始化传入的一致)
	**/
	BizId string `json:"bizId"`
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	 /**
     * 活体检测类型：
         1：远近，
         2：眨眼，
         3：摇头，
         4: 点头，
         5:张嘴，
         6：炫彩，
         7: 读数
     **/
	 LivingType string `json:"livingType"`
	 /**
	 *姓名
	 **/
	 CertName string `json:"certName"`
	 /**
	 *证件号
	 **/
	 CertNo string `json:"certNo"`
	 /**
	 *活体认证截取的最清晰的一张用户照片的下载url地址，有效期为 24 个小时。
	 **/
	 ImgUrl string `json:"imgUrl"`
	/**
     * 认证结果：
         true： 认证成功，
         false：认证失败
    **/
	Pass string `json:"pass"`
}
/**
 * 转换成JSON字符串
 */
 func (self NonMainlandCnVerifyResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}