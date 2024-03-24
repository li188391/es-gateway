package face

import(
	"encoding/json"
)

type DeleteFaceRequest struct{
	/**
	*数据库名称
	**/
	DbName string `json:"dbName"`
	/**
	*实体ID(保持唯一)
	**/
	EntityId string `json:"entityId"`
	
}
/**
*一砂云，请求action
**/
const DeleteFaceRequestAct = "faceId/faceID/deleteFace"
/**
 * 转换成JSON字符串
 */
 func (self DeleteFaceRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}