package util
//连接需要验证密钥，目前设计是统一密钥，没有私钥，后面根据情况调整
// 2018-04-05 by:yq
import (
	"crypto/sha1"
	"time"
)

// 简单的一个校验值，sha加密简单使用
func Getverifyval(verifyKey *string) []byte {
	b := sha1.Sum([]byte(time.Now().Format("2006-01-02 15") + *verifyKey))
	return b[:]
}
