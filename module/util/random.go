package util

import (
	"crypto/rand"
	"math/big"
)

// GetRandomString 生成随机字符串，不为空则成功生成
func GetRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		result = append(result, letters[num.Int64()])
	}

	return string(result)
}
