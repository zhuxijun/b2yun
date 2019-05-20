package util_test

import (
	"io/ioutil"
	"os"
	"b2yun/pkg/util"

	"testing"
)

func TestFileExists(t *testing.T) {

	// 创建临时文件
	file, err := ioutil.TempFile("", "test")
	fileName := file.Name()

	if err != nil {
		t.Error(err)
	}

	// 检测文件存在
	if exits, err := util.FileExist(fileName); !exits || (err != nil) {
		t.Error("文件存在检测错误")
		return
	}

	// 关闭和删除闻临时文件
	file.Close()

	err = os.Remove(fileName)
	if err != nil {
		t.Error(err)
		return
	}

	// 检测文件不存在
	if exits, err := util.FileExist(fileName); (err != nil) || exits {
		t.Error("文件不存在检测错误")
		return
	}

}

func TestReverseString(t *testing.T) {

	s := "123456789asdfghjklqwertyuiopzxcvbnm"

	rS := util.ReverseString(s)
	rrS := util.ReverseString(rS)

	if rrS != s {
		t.Errorf("反转字符串预期不符，期待[%s], 实际[%s]", s, rrS)
	}

	s = "qwerty"
	rS = util.ReverseString(s)
	if rS != "ytrewq" {
		t.Errorf("反转字符串预期不符，期待[%s], 实际[%s]", "ytrewq", rS)
	}
}
