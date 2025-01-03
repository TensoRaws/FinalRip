package ffmpeg

import (
	"os/exec"

	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/util"
)

// MergeVideo 使用 mkvmerge 合并视频，ffmpeg 合并音频和字幕
func MergeVideo(originFile string, inputFiles []string, outputPath string) error {
	tempVideoConcatOutputPath := "temp_video_concat_output.mkv"
	tempVideoMergedOutputPath := "temp_video_merged_output.mkv"
	tempVideoOutputPath := "temp_video_output.mkv"

	// 清理临时文件
	_ = util.ClaerTempFile(tempVideoConcatOutputPath, tempVideoMergedOutputPath, tempVideoOutputPath)
	defer func(p ...string) {
		log.Logger.Infof("Clear temp file %v", p)
		_ = util.ClaerTempFile(p...)
	}(tempVideoConcatOutputPath, tempVideoMergedOutputPath, tempVideoOutputPath)

	// Concat video
	log.Logger.Infof("Concat video with encoded clips: %s", inputFiles)

	mkvmergeArgs := []string{"-o", tempVideoConcatOutputPath}
	for i, file := range inputFiles {
		if i > 0 {
			mkvmergeArgs = append(mkvmergeArgs, "+")
		}
		mkvmergeArgs = append(mkvmergeArgs, file)
	}

	cmd := exec.Command("mkvmerge", mkvmergeArgs...)

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
		tempVideoMergedOutputPath,
	)
	out, err = cmd.CombinedOutput()
	log.Logger.Infof("Merged output: %s", out)
	if err != nil {
		// 清理可能存在的临时文件
		_ = util.ClaerTempFile(tempVideoMergedOutputPath)
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
			tempVideoMergedOutputPath,
		)
		out, err = cmd.CombinedOutput()
		log.Logger.Infof("Merged output: %s", out)
		if err != nil {
			log.Logger.Errorf("Merge audio failed: %v", err)
			return err
		}
	}

	// use mkvmerge to re-mux
	log.Logger.Infof("Re-mux video with mkvmerge and remove tags with mkvpropedit")
	// !mkvmerge -o output.mkv temp_merged.mkv
	// !mkvpropedit output.mkv --tags all:
	cmd = exec.Command(
		"mkvmerge",
		"-o", tempVideoOutputPath,
		tempVideoMergedOutputPath,
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Logger.Errorf("mkvmerge Re-mux video failed: %v", err)
		return err
	}
	log.Logger.Infof("mkvmerge Re-mux output: %s", out)

	// use ffmpeg to re-mux
	cmd = exec.Command(
		"ffmpeg",
		"-i", tempVideoMergedOutputPath,
		"-c", "copy",
		outputPath,
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Logger.Errorf("ffmpeg Re-mux video failed: %v", err)
		return err
	}
	log.Logger.Infof("ffmpeg Re-mux output: %s", out)

	// use mkvpropedit to remove tags
	cmd = exec.Command(
		"mkvpropedit",
		outputPath,
		"--tags", "all:",
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Logger.Errorf("mkvmerge remove tags failed: %v", err)
		return err
	}
	log.Logger.Infof("mkvmerge remove tags output: %s", out)
	return nil
}
