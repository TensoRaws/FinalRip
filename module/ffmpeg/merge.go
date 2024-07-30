package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/util"
)

// MergeVideo 使用 ffmpeg 进行视频合并
func MergeVideo(originFile string, inputFiles []string, outputPath string) error {
	// 写入文件列表
	listPath := "temp_list.txt"
	tempVideoOutputPath := "temp_video_output.mkv"

	defer func(p ...string) {
		log.Logger.Infof("Clear temp file %v", p)
		_ = util.ClaerTempFile(p...)
	}(listPath, tempVideoOutputPath)

	var listStr string
	for _, file := range inputFiles {
		listStr += fmt.Sprintf("file '%s'\n", file)
	}

	err := os.WriteFile(listPath, []byte(listStr), 0755)
	if err != nil {
		log.Logger.Errorf("write list file failed: %v", err)
		return err
	}

	// 执行合并
	commandStr := fmt.Sprintf("ffmpeg -safe 0 -f concat -i %s -c copy %s", listPath, tempVideoOutputPath)
	log.Logger.Infof("Merge video command: %s", commandStr)
	cmd := exec.Command("ffmpeg", "-safe", "0", "-f", "concat", "-i", listPath, "-c", "copy", tempVideoOutputPath) //nolint: lll
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Logger.Errorf("Merge video failed: %v", err)
		return err
	}
	log.Logger.Infof("Merge video output: %s", out)

	// 拼接音频
	commandStr = fmt.Sprintf("ffmpeg -i %s -i %s -map 0:a:0 -map 1:v:0 -c copy %s", originFile, tempVideoOutputPath, outputPath) //nolint: lll
	log.Logger.Infof("Merge audio command: %s", commandStr)
	cmd = exec.Command("ffmpeg", "-i", originFile, "-i", tempVideoOutputPath, "-map", "0:a:0", "-map", "1:v:0", "-c", "copy", outputPath) //nolint: lll
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Logger.Errorf("Merge audio failed: %v", err)
		return err
	}
	log.Logger.Infof("Merge audio output: %s", out)

	return nil
}
