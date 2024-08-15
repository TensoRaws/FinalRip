package util

import (
	"strconv"
)

func GenerateClipKey(k string, i int) string {
	return k + "-clip-" + strconv.FormatInt(int64(i), 10) + ".mkv"
}

func GenerateClipEncodedKey(k string, i int) string {
	return k + "-clip-encoded-" + strconv.FormatInt(int64(i), 10) + ".mkv"
}

func GenerateMergedKey(k string) string {
	return k + "-Encoded" + ".mkv"
}
