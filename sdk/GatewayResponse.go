package sdk

import (
	"encoding/json"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"log"
	
)

type GatewayResponse struct {
	/**
	 * 协议版本
	 */
	Ver string
	/**
	 * 响应状态码
	 */
	Code string
	/**
	 * 结果描述
	 */
	Msg string
	/**
	 *
	 */
	GwRequestID string
	/**
	 * 业务action
	 */
	Act string
	/**
	 * 响应时间戳
	 */
	Timestamp string
	/**
	 * 业务响应数据
	 */
	BizContent string
	/**
	 * 签名数据
	 */
	Sign string
	/**
	 * 签名类型
	 */
	SignType string
	/**
	 * 验签结果
	 */
	SignVerifyResult bool
}

func (self GatewayResponse) initWithCode(code string, msg string) {
	self.Code = code
	self.Msg = msg
	if code == "-1" {
		fmt.Println("异常 :", msg)
	}
}

func (self *GatewayResponse) verifySign(key string) {
	// TODU 对数据进行验签
	switch self.SignType {
	case "1":
	case "2":
	case "3":
		var builder strings.Builder
		builder.WriteString(self.Act)
		builder.WriteString("&")
		builder.WriteString(self.Timestamp)
		builder.WriteString("&")
		builder.WriteString(self.Code)
		if self.BizContent != "" {
			builder.WriteString("&")
			builder.WriteString(self.BizContent)
		}
		builder.WriteString("&")
		builder.WriteString(key)
		signedData := []byte(builder.String())
		log.Printf("验签数据: %s", string(signedData))
		hash := md5.New()
		hash.Write(signedData)
		signBuff := hash.Sum(nil)
		signHexLowerCase := hex.EncodeToString(signBuff)
		if strings.ToLower(signHexLowerCase) == strings.ToLower(self.Sign) { 
			self.SignVerifyResult = true
		}
	default:
	}

}

func (self GatewayResponse) ToJsonStr() string {
	jsonStr, _ := json.Marshal(self)

	return string(jsonStr)
}
