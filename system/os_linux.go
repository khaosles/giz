//go:build linux

package system

import (
	"os/exec"
	"syscall"
)

/*
   @File: os_linux.go
   @Author: khaosles
   @Time: 2023/8/13 13:17
   @Desc:
*/

func WithForeground() Option {
	return func(c *exec.Cmd) {
		if c.SysProcAttr == nil {
			c.SysProcAttr = &syscall.SysProcAttr{
				Foreground: true,
			}
		} else {
			c.SysProcAttr.Foreground = true
		}
	}
}

func WithWinHide() Option {
	return func(c *exec.Cmd) {

	}
}
