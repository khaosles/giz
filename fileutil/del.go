package fileutil

import (
	"os"
	"path/filepath"

	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: del.go
   @Author: khaosles
   @Time: 2023/8/23 15:35
   @Desc:
*/

// Delete 删除文件并删除{tier}层空文件夹
func Delete(path string, tier int) {
	if tier < 0 {
		tier = 0
	}
	err := os.Remove(path)
	if err != nil {
		glog.Error(err)
		return
	}
	for i := 0; i < tier; i++ {
		path = Dirname(path)
		result := removeEmptyDir(path)
		if !result {
			break
		}
	}
	return
}

func removeEmptyDir(dir string) bool {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		glog.Error(err)
		return false
	}
	if len(files) == 0 {
		err := os.Remove(dir)
		if err != nil {
			glog.Error(err)
			return false
		}
		return true
	}
	return false
}
