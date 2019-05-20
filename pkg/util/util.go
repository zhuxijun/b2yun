package util

import "os"

// FileExist 判断文件是否存在，文件读取失败否会error
func FileExist(path string) (bool, error) {

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

// ReverseString 反转字符串
func ReverseString(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
