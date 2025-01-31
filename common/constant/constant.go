package constant

import "strings"

const (
	FINALRIP                  = "FINALRIP"
	ENV_FINALRIP_SOURCE       = FINALRIP + "_SOURCE"
	FINALRIP_SOURCE_MKV       = FINALRIP + "_SOURCE.mkv"
	FINALRIP_ENCODED_CLIP     = FINALRIP + "_ENCODED_CLIP"
	FINALRIP_ENCODED_CLIP_MKV = FINALRIP + "_ENCODED_CLIP.mkv"
)

// ContainsFinalRipInString 检查字符串中是否包含 FinalRip 常量
func ContainsFinalRipInString(s string, constant string) bool {
	// 如果待检查的字符串中不包含 FINALRIP 字符，则不是合法的 FinalRip 常量
	if !strings.Contains(constant, FINALRIP) {
		return false
	}

	return strings.Contains(s, constant)
}
