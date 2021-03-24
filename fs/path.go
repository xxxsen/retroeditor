package fs

import (
	"strings"
)

//MergePath 将基础路径跟相对路径合并成一个绝对路径
func MergePath(parent string, child string) string {
	p := SplitPath(parent)
	c := SplitPath(child)
	p = append(p, c...)
	return "/" + strings.Join(p, "/")
}

//TrimPath 从子路径中移除父路径, 生成相对路径
func TrimPath(parent string, child string) string {
	p := SplitPath(parent)
	c := SplitPath(child)
	begin := 0
	for i := 0; i < len(p) && i < len(c); i++ {
		if p[i] == c[i] {
			begin = i + 1
		} else {
			begin = i
			break
		}
	}
	diff := c[begin:]
	return strings.Join(diff, "/")
}

//SplitPath 将路径进行切割
func SplitPath(path string) []string {
	var lst []string
	last := 0
	var i = 0
	for i = 0; i < len(path); i++ {
		if path[i] == '/' || path[i] == '\\' {
			lst = append(lst, path[last:i])
			last = i + 1
		}
	}
	if i != last {
		lst = append(lst, path[last:i])
	}
	var rs []string
	for _, item := range lst {
		if item == "." || item == "" {
			continue
		}
		if item == ".." {
			if len(rs) > 0 {
				rs = rs[0 : len(rs)-1]
			}
			continue
		}
		rs = append(rs, item)
	}
	if len(rs) == 0 {
		return nil
	}
	return rs
}

func IsParentDir(parent string, child string) bool {
	parentLst := SplitPath(parent)
	childLst := SplitPath(child)
	if len(parentLst) == 0 {
		return true
	}
	if len(parentLst) > len(childLst) {
		return false
	}
	for i := 0; i < len(parentLst); i++ {
		if parentLst[i] != childLst[i] {
			return false
		}
	}
	return true
}
