package http_test

import (
	"errors"
	"testing"
	root "b2yun/pkg"
	"b2yun/pkg/http"
)

type Config struct {
	Host  string
	Error bool
}

func (c *Config) GetConfig() (*root.Config, error) {
	if c.Error == false {
		return &root.Config{
			HTTP: &root.HTTPConfig{
				Host: c.Host,
			},
		}, nil
	} else {
		return nil, errors.New("server_config_error")
	}

}
func TestServer(t *testing.T) {

	c := Config{
		Host: ":2222",
	}
	handler := http.NewHandler(root.NewLogStdOut())

	server := http.NewServer(&c, handler)

	go server.Open()
	go server.Open()

	defer server.Close()

}

func TestServer_NilHost(t *testing.T) {

	c := Config{}

	handler := http.NewHandler(root.NewLogStdOut())

	server := http.NewServer(&c, handler)

	err := server.Open()
	if root.ErrorCode(err) != "config_http_not_found" {
		t.Error(err)
	}

}

func TestServer_OpenConfigError(t *testing.T) {

	c := Config{Error: true}
	handler := http.NewHandler(root.NewLogStdOut())

	server := http.NewServer(&c, handler)

	err := server.Open()
	if root.ErrorCode(err) != "http_server_open_getconfig_err" {
		t.Error(err)
	}
}
