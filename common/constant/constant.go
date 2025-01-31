package constant

import "strings"

type FinalRipConstant string

const (
	FINALRIP                  = FinalRipConstant("FINALRIP")
	ENV_FINALRIP_SOURCE       = FINALRIP + FinalRipConstant("_SOURCE")
	FINALRIP_SOURCE_MKV       = FINALRIP + FinalRipConstant("_SOURCE.mkv")
	FINALRIP_ENCODED_CLIP     = FINALRIP + FinalRipConstant("_ENCODED_CLIP")
	FINALRIP_ENCODED_CLIP_MKV = FINALRIP + FinalRipConstant("_ENCODED_CLIP.mkv")
)

// ContainsFinalRipInString 检查字符串中是否包含 FinalRip 常量
func ContainsFinalRipInString(s string, constant FinalRipConstant) bool {
	return strings.Contains(s, string(constant))
}
