package sdk

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type GatewayRequest struct {
	/**
	 * 用户appCode,参考 : https://esandinfo.yuque.com/yv6e1k/aa4qsg/cdwove
	 */
	AppCode string `json:"appCode"`
	/**
	 * 业务action（参考具体业务协议文档）
	 */
	Act string `json:"act"`
	/**
	 * 业务ID (随机数，保证其唯一)
	 */
	BizNo string `json:"bizNo"`
	/**
	 * 当前时间戳(时区为北京时间，防黑客重放攻击)，格式为：yyyyMMddHHmmssSSS
	 */
	Timestamp string `json:"timestamp"`
	/**
	 * 签名类型，1：ECC,2:RSA,3:MD5
	 */
	Type string `json:"type"`
	/**
	 * 业务内容
	 */
	BizContent string `json:"bizContent"`
	/**
	 * 数据签名，可参考下面章节 “业务服务器签名验签”
	 */
	Sign string `json:"sign"`
}

/**
* 对数据进行签名
* @param key    签名密钥
* @return 签名执行结果
 */
func (self *GatewayRequest) genSign(key string) {
	switch self.Type {
	case "1":
	case "2":
	case "3":
		signData := self.BizNo + "&" + self.AppCode + "&" + self.BizContent + "&" + self.Timestamp + "&" + key
		// 算MD5
		md5Sum := md5.Sum([]byte(signData))
		md5SumHex := make([]byte, hex.EncodedLen(len(md5Sum)))
		hex.Encode(md5SumHex, md5Sum[:])
		self.Sign = string(md5SumHex)
	default:
	}

}

func (self GatewayRequest) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
