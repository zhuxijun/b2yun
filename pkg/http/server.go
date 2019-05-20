package http

import (
	root "b2yun/pkg"
	"net"
	"net/http"
)

// Server http服务
type Server struct {
	ln net.Listener

	Handler *Handler

	configer root.Configer

	server *http.Server
}

// NewServer 创建web服务器
func NewServer(config root.Configer, handler *Handler) *Server {
	s := &Server{
		configer: config,
		Handler:  handler,
		server:   &http.Server{},
	}
	return s
}

// Open 打开web服务
func (s *Server) Open() error {

	const op = "http.Server.Open"
	var customError root.Error
	customError.Op = op

	config, err := s.configer.GetConfig()
	if err != nil {
		customError.Code = "http_server_open_getconfig_err"
		customError.Err = err
		return &customError
	}

	if config.HTTP.Host == "" {
		customError.Code = "config_http_not_found"
		return &customError
	}

	s.server.Addr = config.HTTP.Host
	s.server.Handler = s.Handler

	s.server.ListenAndServe()

	return nil
}

// Close 关闭web服务
func (s *Server) Close() error {

	if s.server != nil {
		s.server.Close()
	}
	return nil
}
