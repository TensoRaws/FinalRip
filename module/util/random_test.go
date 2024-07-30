package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomString(t *testing.T) {
	a := assert.New(t)
	a.Equal(10, len(GetRandomString(10)))
	a.Equal(20, len(GetRandomString(20)))
	// 生成100长度字符串，测试是否有重复
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			a.NotEqual(GetRandomString(100), GetRandomString(100))
		}
	}
}
