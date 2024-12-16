package util

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/process"
)

// KillProcessByName 根据进程名杀死进程
func KillProcessByName(processName string) error {
	// 获取所有进程
	processes, err := process.Processes()
	if err != nil {
		return fmt.Errorf("cannot get processes: %v", err)
	}

	// 遍历进程，匹配进程名
	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}

		// 如果进程名匹配，则杀死进程
		if name == processName || name == processName+".exe" {
			err := p.Kill()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
