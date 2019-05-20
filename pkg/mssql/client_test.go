package mssql_test

import (
	"errors"
	root "b2yun/pkg"
	"b2yun/pkg/mssql"
	"testing"
)

// Config 配置对象
type Config struct {
	ConfigFn func() (*root.Config, error)
}

func (c *Config) GetConfig() (*root.Config, error) {
	return c.ConfigFn()
}

// Client 对象，测试用
type Client struct {
	client *mssql.Client
}

func TestClient_OK(t *testing.T) {

	var config Config
	config.ConfigFn = func() (*root.Config, error) {
		rConfig := root.Config{}
		rConfig.Mssql = &root.MssqlConfig{
			Parm: "",
		}
		return &rConfig, nil
	}

	client := mssql.NewClient(&config)
	err := client.Open()

	if err != nil {
		t.Error(err)
	}

	defer client.Close()
}

func TestClient_GetConfigError(t *testing.T) {

	config := Config{}
	config.ConfigFn = func() (*root.Config, error) {
		return nil, errors.New("testGetConfig")
	}

	client := mssql.NewClient(&config)

	if err := client.Open(); root.ErrorCode(err) != mssql.ECONFIGINVALID {
		t.Errorf("配置文件错误判断出错,期待%s, 实际%s", mssql.ECONFIGINVALID, root.ErrorCode(err))
	}
	defer client.Close()

}

func TestClient_MssqlConfigNil(t *testing.T) {

	var config Config
	config.ConfigFn = func() (*root.Config, error) {
		return &root.Config{}, nil
	}

	client := mssql.NewClient(&config)

	if err := client.Open(); root.ErrorCode(err) != mssql.ECONFIGMSSQLNOTFOUND {
		t.Errorf("错误返回值不符合预期，预期[%s]，实际[%s]", mssql.ECONFIGMSSQLNOTFOUND, root.ErrorCode(err))
	}

	defer client.Close()
}

func TestConfig_OpenError(t *testing.T) {
	var config Config
	config.ConfigFn = func() (*root.Config, error) {
		rConfig := root.Config{}
		rConfig.Mssql = &root.MssqlConfig{
			Parm:     "testOpen",
			Dialects: "not_exists_dialects",
		}
		return &rConfig, nil
	}

	client := mssql.NewClient(&config)

	if err := client.Open(); root.ErrorCode(err) != mssql.EDBOPENERROR {
		t.Errorf("错误返回值不符合预期，预期[%s], 实际[%s]", mssql.EDBOPENERROR, root.ErrorCode(err))
	}

	defer client.Close()
}

func TestClient_Connect(t *testing.T) {
	var config Config
	client := mssql.NewClient(&config)
	client.Connect()
	defer client.Close()

}
