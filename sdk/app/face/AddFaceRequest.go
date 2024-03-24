package face

import(
	"encoding/json"
)

type AddFaceRequest struct{
	/**
	*数据库名称
	**/
	DbName string `json:"dbName"`
	/**
	*实体ID(保持唯一)
	**/
	EntityId string `json:"entityId"`
	/**
	*人脸照片B64数据，
	**/
	Image string `json:"image"`
	/**
	*自定义信息。支持字母、数字、标点符号和汉字。不超过512个字符
	**/
	ExtraData string `json:"extraData"`
}
/**
*一砂云，请求action
**/
const AddFaceRequestAct = "faceId/faceID/addFace"
/**
 * 转换成JSON字符串
 */
 func (self AddFaceRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}