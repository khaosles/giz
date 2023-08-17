package g

import (
	"os"
	"os/signal"
	"syscall"
)

/*
   @File: exit.go
   @Author: khaosles
   @Time: 2023/8/17 18:01
   @Desc:
*/

func Exit(cb func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	s := <-ch
	cb()
	if i, ok := s.(syscall.Signal); ok {
		os.Exit(int(i))
	} else {
		os.Exit(0)
	}
}
