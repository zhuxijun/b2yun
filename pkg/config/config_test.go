package config_test

import (
	root "b2yun/pkg"
	"b2yun/pkg/config"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

// TestFileNotFound 测试文件不存在
func TestFileNotFound(t *testing.T) {
	path := "not found path"
	c := config.NewConfig(path)

	_, err := c.GetConfig()
	if root.ErrorCode(err) != config.ECONFIGNOTFOUND {
		t.Error(err)
	}
}

//TestConfNotInValid 测试配置文件json格式不合法
func TestConfNotInValid(t *testing.T) {

	tmpFile, err := ioutil.TempFile("", "app_config.json")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	notInvalidContent := "a: 1"
	tmpFile.WriteString(notInvalidContent)

	path := tmpFile.Name()

	c := config.NewConfig(path)

	_, err = c.GetConfig()
	if root.ErrorCode(err) != config.ECONFIGINVALID {
		t.Error(err)
	}

}

//TestConfInvalid 测试能正确解析配置
func TestConfInvalid(t *testing.T) {

	path, err := getConfigWithTempPath()
	if err != nil {
		t.Error(err)
	}

	out, err := getConfigJsonInvalid()

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 666)

	defer os.Remove(file.Name())
	defer file.Close()

	if err != nil {
		t.Error(err)
	}
	_, err = file.Write(out)
	if err != nil {
		t.Error(err)
	}

	config_out := config.NewConfig(path)

	_, err = config_out.GetConfig()
	if err != nil {
		t.Error(err)
	}

}

// GetConfigWithTempPath 从空文件中获取配置信息
func getConfigWithTempPath() (string, error) {
	tmpFile, err := ioutil.TempFile("", "app_config.json")
	defer tmpFile.Close()
	if err != nil {
		return tmpFile.Name(), err
	}
	return tmpFile.Name(), nil

}

func getConfigJsonInvalid() ([]byte, error) {
	mssqlPath := root.MssqlConfig{
		Parm: "sqlserver://sa:123@localhost:1433?database=zbhb",
	}

	config_in := root.Config{
		Mssql: &mssqlPath,
	}

	return json.Marshal(config_in)

}
