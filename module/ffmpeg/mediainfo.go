package ffmpeg

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type VideoInfo struct {
	FrameCount int     `json:"frame_count"`
	Duration   float64 `json:"duration"`
	Resolution string  `json:"resolution"`
}

// GetVideoInfo 使用 ffprobe 提取视频信息
func GetVideoInfo(filePath string) (VideoInfo, error) {
	var info VideoInfo

	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0",
		"-show_entries", "stream=nb_read_frames:duration:width:height",
		"-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return info, err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) < 4 {
		return info, fmt.Errorf("unexpected output from ffprobe")
	}

	frameCount, err := strconv.Atoi(lines[0])
	if err != nil {
		return info, err
	}
	duration, err := strconv.ParseFloat(lines[1], 64)
	if err != nil {
		return info, err
	}
	resolution := fmt.Sprintf("%s x %s", lines[2], lines[3])

	info = VideoInfo{
		FrameCount: frameCount,
		Duration:   duration,
		Resolution: resolution,
	}

	return info, nil
}
