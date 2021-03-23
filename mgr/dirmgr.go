package mgr

import (
	"io/ioutil"
)

//DirManager 目录管理
type DirManager struct {
	base string
}

//NewDirManager 新建
func NewDirManager(base string) *DirManager {
	return &DirManager{base: base}
}

//GetDirs 获取当前目录下的所有目录
func (d *DirManager) GetDirs() ([]string, error) {
	arr, err := ioutil.ReadDir(d.base)
	if err != nil {
		return nil, err
	}
	var fs []string
	for _, item := range arr {
		if !item.IsDir() {
			continue
		}
		fs = append(fs, item.Name())
	}
	return fs, nil
}

//GetBase 获取当前的基础地址
func (d *DirManager) GetBase() string {
	return d.base
}
