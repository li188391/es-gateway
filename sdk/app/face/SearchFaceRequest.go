package face

import(
	"encoding/json"
)

type SearchFaceRequest struct{
	/**
	*数据库名称
	**/
	DbName string `json:"dbName"`
	/**
	*人脸照片B64数据，
	**/
	Image string `json:"image"`
	/**
	*是否开启活体检测（默认false）
	**/
	EnableLiving bool `json:"enableLiving"`
}
/**
*一砂云，请求action
**/
const SearchFaceRequestAct = "faceId/faceID/searchFace"
/**
 * 转换成JSON字符串
 */
 func (self SearchFaceRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}