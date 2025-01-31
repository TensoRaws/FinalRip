package constant

import (
	"errors"
	"strings"
	"unicode"
)

type FinalRipConstant string

const (
	FINALRIP                  = FinalRipConstant("FINALRIP")
	ENV_FINALRIP_SOURCE       = FINALRIP + FinalRipConstant("_SOURCE")
	FINALRIP_SOURCE_MKV       = FINALRIP + FinalRipConstant("_SOURCE.mkv")
	FINALRIP_ENCODED_CLIP     = FINALRIP + FinalRipConstant("_ENCODED_CLIP")
	FINALRIP_ENCODED_CLIP_MKV = FINALRIP + FinalRipConstant("_ENCODED_CLIP.mkv")
)

// CheckVSScriptAndEncodeParam 检查传入的 Script 和 EncodeParam 是否合法
func CheckVSScriptAndEncodeParam(script string, encodeParam string) error {
	if !strings.Contains(script, string(ENV_FINALRIP_SOURCE)) {
		return errors.New("the VapourSynth Script code must contain " + string(ENV_FINALRIP_SOURCE) + " environment variable to specify the source video") //nolint:lll
	}
	if !strings.Contains(encodeParam, string(FINALRIP_ENCODED_CLIP_MKV)) {
		return errors.New("the Encode Param must contain " + string(FINALRIP_ENCODED_CLIP_MKV) + " to specify the output video clip") //nolint:lll
	}
	encodeParam = strings.TrimRightFunc(encodeParam, unicode.IsSpace) // 去除末尾空格，换行符等
	if strings.Contains(encodeParam, "\n") || strings.Contains(encodeParam, "\r") {
		return errors.New("the Encode Param cannot contain line break")
	}
	return nil
}
