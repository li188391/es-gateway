package face

import(
	"encoding/json"
)

type CreateFaceDbRequest struct{
	/**
	*数据库名称
	**/
	DbName string `json:"dbName"`
	/**
	*数据库名称
	**/
	Description string `json:"description"`
}
/**
*一砂云，请求action
**/
const CreateFaceDbRequestAct = "faceId/faceID/createFaceDb"
/**
 * 转换成JSON字符串
 */
 func (self CreateFaceDbRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}