package mgr

import (
	"io/fs"
	"path/filepath"
)

type FileMgr struct {
	base string
}

//NewFileMgr NewFileMgr
func NewFileMgr(base string) *FileMgr {
	m := &FileMgr{base: base}
	return m
}

//Search Search
func (f *FileMgr) Search(filters map[string]bool) ([]string, error) {
	var rs []string
	err := filepath.Walk(f.base, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(info.Name())
		if _, ok := filters[ext]; ok {
			rs = append(rs, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return rs, nil
}
