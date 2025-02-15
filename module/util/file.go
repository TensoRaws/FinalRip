package util

import (
	"math/big"
	"os"

	"github.com/dustin/go-humanize"
)

// ByteCountBinary 把 Length 转换为文件大小, 自适应单位
func ByteCountBinary(b uint64) string {
	// uint64 to bigint
	bigInt := new(big.Int).SetUint64(b)
	return humanize.BigIBytes(bigInt)
}

// ClearTempFile 清理临时文件
func ClearTempFile(tempPath ...string) error {
	for _, p := range tempPath {
		err := os.RemoveAll(p)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFileSize 获取文件大小
func GetFileSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}
