package livingdetect

import(
	"encoding/json"
)

type LivingdetectVerifyResponse struct{
	/**
     * 错误码
    **/
	Code string `json:"code"`
	/**
	*结果描述
	**/
	Msg string `json:"msg"`
	/**
	*请求ID
	**/
	RequestId string `json:"requestId"`
	/**
	*活体认证截取的最清晰的一张用户照片的下载url地址，有效期为 24 个小时
	*/
	BestImg string `json:"bestImg"`
	/**
	*如果是需要多个照片返回时，不会返回bestImg，转而返回这个字段，一个数组
	**/
	BestImgArray string `json:"bestImgArray"`
	/**
	*如果是需要多个照片返回时，前端不想要照片上传到oss时，会返回这个字段
	**/
	BestImgBase64Array string `json:"bestImgBase64Array"`
	/**
	*活体检测状态码，LDT_SUCCESS： 活体检测成功，LDT_FAILED：活体检测失败
	**/
	ResultCode string `json:"resultCode"`	
}
/**
 * 转换成JSON字符串
 */
 func (self LivingdetectVerifyResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}