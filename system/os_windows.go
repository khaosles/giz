//go:build windows

package system

/*
   @File: os_windows.go
   @Author: khaosles
   @Time: 2023/8/13 13:16
   @Desc:
*/

func WithWinHide() Option {
	return func(c *exec.Cmd) {
		if c.SysProcAttr == nil {
			c.SysProcAttr = &syscall.SysProcAttr{
				HideWindow: true,
			}
		} else {
			c.SysProcAttr.HideWindow = true
		}
	}
}

func WithForeground() Option {
	return func(c *exec.Cmd) {

	}
}
