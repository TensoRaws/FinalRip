package ffmpeg

import (
	"bufio"
	"fmt"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/util"
	"io"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

// EncodeVideo 压制视频，压制后的视频文件名为 encoded.mkv，压制参数由 encodeParam 指定，压制视频从环境变量 FINALRIP_SOURCE 读取
func EncodeVideo(encodeScript string, encodeParam string) error {
	encodeScriptPath := "encode.py"
	// encodedVideo := "encoded.mkv"
	// 根据操作系统创建脚本文件
	var commandStr string
	var scriptPath string
	var condaInitScript string
	switch runtime.GOOS {
	case "windows":
		scriptPath = "temp_script.bat"
		condaInitScript = "@echo off\r\n" + "call \"%USERPROFILE%\\miniconda3\\condabin\\activate.bat\"\r\n"
	default:
		scriptPath = "temp_script.sh"
		condaInitScript = "#!/bin/bash\n" + "source ~/miniconda3/etc/profile.d/conda.sh\n" + "conda activate\n"
	}
	commandStr = condaInitScript + "vspipe -c y4m encode.py - | ffmpeg -i - encoded.mkv " + encodeParam
	log.Logger.Info("commandStr: " + commandStr)

	// 清理临时文件
	_ = util.ClaerTempFile(encodeScriptPath, scriptPath)

	// 写入压制 py
	err := os.WriteFile(encodeScriptPath, []byte(encodeScript), 0755)
	if err != nil {
		log.Logger.Error("write vapoursynth script file failed: " + err.Error())
		return err
	}
	// 写入脚本文件
	err = os.WriteFile(scriptPath, []byte(commandStr), 0755)
	if err != nil {
		log.Logger.Error("write script file failed: " + err.Error())
		return err
	}
	// 执行脚本
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", scriptPath)
	} else {
		cmd = exec.Command("sh", scriptPath)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Logger.Error("error: " + err.Error())
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Logger.Error("error: " + err.Error())
	}

	err = cmd.Start()
	if err != nil {
		log.Logger.Error("cmd start error: " + err.Error())
		return err
	}

	var wg sync.WaitGroup
	readerFunc := func(pipe io.Reader) {
		defer wg.Done()
		scanner := bufio.NewScanner(pipe)
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Logger.Error("error reading pipe: " + err.Error())
		}
	}
	wg.Add(2)
	go readerFunc(stdout)
	go readerFunc(stderr)

	err = cmd.Wait()
	if err != nil {
		log.Logger.Error("cmd wait error: " + err.Error())
		return err
	}
	wg.Wait()

	// 清理临时文件
	err = util.ClaerTempFile(encodeScriptPath, scriptPath)
	if err != nil {
		log.Logger.Error("clear temp file failed: " + err.Error())
		return err
	}

	return nil
}
