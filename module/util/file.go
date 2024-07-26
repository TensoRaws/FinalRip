package util

import (
	"math/big"

	"github.com/dustin/go-humanize"
)

// ByteCountBinary 把 Length 转换为文件大小, 自适应单位
func ByteCountBinary(b uint64) string {
	// uint64 to bigint
	bigInt := new(big.Int).SetUint64(b)
	return humanize.BigIBytes(bigInt)
}
