package pty

import "bytes"

// 判断切片x 是否包含y ，包含则返回去除y后的x，否则返回false
func IsContain(x, y []byte) (n []byte, contain bool) {
	index := bytes.Index(x, y)
	if index == -1 {
		return n, false
	}
	lastIndex := index + len(y)
	n = append(x[:index], x[lastIndex:]...)
	return n, true
}
