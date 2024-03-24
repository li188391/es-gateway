package rpc

import(
	"encoding/json"
)

type VerifyRequest struct{
	/**
	*姓名
	**/
	CertName string `json:"certName"`
	/**
	*身份证号码
	**/
	CertNo string `json:"certNo"`
	/**
	*获取活体检测行为”接口返回的活体行为特征唯一标识
	**/
	BehaviorToken string `json:"behaviorToken"`
	/**
	*包含人脸以及活体检测行为的短视频（Base64URLSafe编码串，注：设置该字段时，videoUrl字段设置为空。要求为ffmpeg所支持的格式及码率，建议视频及音频采样率8000以上，帧速率二三十帧/秒即可，分辨率720P，视频时长为3~6秒，视频的大小不能大于8M，过大的视频分辨率会导致视频处理时间比较长，容易导致超时问题）
    **/
	Video string `json:"video"`
	/**
	*录制的活体读数字人脸视频在互联网上的下载链接(设置该字段时，video字段设置为空,若下载地址链接包含&符号，需要使用%26转义符替代)
	**/
	VideoUrl string `json:"videoUrl"`
	
} 
/**
*阿里云请求URL
**/
const VerifyRequestALIYUN_URL = "https://edis3v.market.alicloudapi.com/verify"
/**
 * 转换成JSON字符串
 */
 func (self VerifyRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
