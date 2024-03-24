package rpverifyExt

import (
	"encoding/json"
)

type UploadIDIMGRequest struct{
	/**
	*业务ID(必须保证唯一，方便后续业务跟踪)
	**/
	BizId string `json:"bizId"`
	/**
	*证件照片 (BASE64字符串)，必须是人像面
	**/
	IDImg string `json:"IDImg"`
	/**
	*国家 ISO CODE (可参考 https://esandinfo.yuque.com/yv6e1k/aa4qsg/lu0ya9fog51nsq8l)
	*/
	CountryISOCode string `json:"countryISOCode"`
	/**
	*证件类型 ：
    *passport : 护照
    *identityCard : 身份证 
	**/
	IDType string `json:"IDType"`
	/**
	*证件号
	**/
	IDNumber string `json:"IDNumber"`
	/**
	*姓名
	**/
	Name string `json:"name"`
}

/**
* 一砂云对应的ACTION值
**/
const UploadIDIMGRequestAct = "livingdetection/rpverifyExt/uploadIDIMG"
/**
 * 转换成JSON字符串
 */
func (self UploadIDIMGRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
