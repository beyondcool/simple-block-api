package utils

import (
	"path/filepath"
	"runtime"
)

func GetCurrentGoFileDir() string {
	// 关键：0 代表当前函数自身，1 代表调用者，这里固定填 1
	_, file, _, _ := runtime.Caller(1)

	// 取文件的目录
	dir := filepath.Dir(file)

	return dir
}
