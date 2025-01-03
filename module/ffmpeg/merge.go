package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/util"
)

// MergeVideo 使用 ffmpeg 进行视频合并，使用 mkvpropedit 清除 tags
func MergeVideo(originFile string, inputFiles []string, outputPath string) error {
	listPath := "temp_list.txt" // 写入文件列表
	tempVideoConcatOutputPath := "temp_video_concat_output.mkv"

	// clear temp file
	_ = util.ClaerTempFile(listPath, tempVideoConcatOutputPath)
	defer func(p ...string) {
		log.Logger.Infof("Clear temp file %v", p)
		_ = util.ClaerTempFile(p...)
	}(listPath, tempVideoConcatOutputPath)

	var listStr string
	for _, file := range inputFiles {
		listStr += fmt.Sprintf("file '%s'\n", file)
	}

	err := os.WriteFile(listPath, []byte(listStr), 0755)
	if err != nil {
		log.Logger.Errorf("write list file failed: %v", err)
		return err
	}

	// Concat video
	log.Logger.Infof("Concat video with list: %s", listPath)
	cmd := exec.Command(
		"ffmpeg",
		"-safe", "0",
		"-f", "concat",
		"-i", listPath,
		"-c", "copy",
		tempVideoConcatOutputPath,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Logger.Errorf("Concat video failed: %v", err)
		return err
	}
	log.Logger.Infof("Concat video output: %s", out)

	// Merge video
	log.Logger.Infof("Merge video with concat output: %s", tempVideoConcatOutputPath)
	// audio track + subtitle track
	cmd = exec.Command(
		"ffmpeg",
		"-i", originFile,
		"-i", tempVideoConcatOutputPath,
		"-map", "1:v:0",
		"-map", "0:a",
		"-map", "0:s",
		"-disposition:v:0", "default",
		"-c:v", "copy",
		"-c:a", "flac",
		"-c:s", "copy",
		"-max_interleave_delta", "0",
		outputPath,
	)
	out, err = cmd.CombinedOutput()
	log.Logger.Infof("Merged output: %s", out)
	if err != nil {
		// clean maybe failed output
		_ = util.ClaerTempFile(outputPath)
		log.Logger.Errorf("Merge audio with audio and subtitle failed: %v, try to merge audio only", err)
		// audio track
		cmd = exec.Command(
			"ffmpeg",
			"-i", originFile,
			"-i", tempVideoConcatOutputPath,
			"-map", "1:v:0",
			"-map", "0:a",
			"-disposition:v:0", "default",
			"-c:v", "copy",
			"-c:a", "flac",
			"-max_interleave_delta", "0",
			outputPath,
		)
		out, err = cmd.CombinedOutput()
		log.Logger.Infof("Merged output: %s", out)
		if err != nil {
			log.Logger.Errorf("Merge audio failed: %v", err)
			return err
		}
	}

	// Remove tags with mkvpropedit
	log.Logger.Infof("Remove tags with mkvpropedit...")
	cmd = exec.Command(
		"mkvpropedit",
		outputPath,
		"--tags", "all:",
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Logger.Errorf("Remove tags failed: %v", err)
		return err
	}
	log.Logger.Infof("Remove tags output: %s", out)
	return nil
}
