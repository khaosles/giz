//go:build darwin

package system

import "os/exec"

/*
   @File: os_darwin.go
   @Author: khaosles
   @Time: 2023/8/13 13:16
   @Desc:
*/

func WithForeground() Option {
	return func(c *exec.Cmd) {

	}
}

func WithWinHide() Option {
	return func(c *exec.Cmd) {

	}
}
