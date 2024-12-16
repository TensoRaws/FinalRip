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
	tempVideoConcatOutputPath := "temp_video_concat_output.mkv"
	tempVideoMergedOutputPath := "temp_video_merged_output.mkv"

	// 清理临时文件
	_ = util.ClaerTempFile(listPath, tempVideoConcatOutputPath, tempVideoMergedOutputPath)
	defer func(p ...string) {
		log.Logger.Infof("Clear temp file %v", p)
		_ = util.ClaerTempFile(p...)
	}(listPath, tempVideoConcatOutputPath, tempVideoMergedOutputPath)

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
		"-c", "copy",
		tempVideoMergedOutputPath,
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Logger.Errorf("Merge audio with audio and subtitle failed: %v, try to merge audio only", err)
		// audio track
		cmd = exec.Command(
			"ffmpeg",
			"-i", originFile,
			"-i", tempVideoConcatOutputPath,
			"-map", "1:v:0",
			"-map", "0:a",
			"-disposition:v:0", "default",
			"-c", "copy",
			tempVideoMergedOutputPath,
		)
		out, err = cmd.CombinedOutput()
		if err != nil {
			log.Logger.Errorf("Merge audio failed: %v", err)
			return err
		}
	}
	log.Logger.Infof("Merged output: %s", out)

	// 使用 mkvtoolnix 删除多余的 tags，重新混流
	log.Logger.Infof("Re-mux video with mkvmerge and remove tags with mkvpropedit")
	// !mkvmerge -o output.mkv temp_merged.mkv
	// !mkvpropedit output.mkv --tags all:
	cmd = exec.Command(
		"mkvmerge",
		"-o", outputPath,
		tempVideoMergedOutputPath,
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Logger.Errorf("Re-mux video failed: %v", err)
		return err
	}
	log.Logger.Infof("Re-mux output: %s", out)
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
