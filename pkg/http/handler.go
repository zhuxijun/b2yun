package http

import (
	"encoding/json"
	root "b2yun/pkg"
	"net/http"
)

// Handler 处理器
type Handler struct {
	Log *root.Log
}

// Services 服务集
type Services struct{}

// NewHandler 创建新的处理器
func NewHandler(log *root.Log) *Handler {
	h := &Handler{Log: log}
	return h
}

// Init 初始化处理器函数
func (h *Handler) Init(s Services) {}

// ServeHTTP 开启http服务
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// Error api错误处理
func Error(w http.ResponseWriter, err error, code int, log *root.Log) {

	// 如果错误码 not_found则跳转到not_found
	if root.ErrorCode(err) == root.ENOFOUND {
		NotFound(w)
		return
	}

	customErr := root.Error{}
	// 记录错误日志
	log.Logger.Error(err)
	// 隐藏服务器内部错误
	if code == http.StatusInternalServerError {
		customErr.Code = root.EINTERNAL
	} else {
		customErr.Err = err
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&errorResponse{Err: root.ErrorMessage(&customErr)})
}

// errorResponse 通用错误返回
type errorResponse struct {
	Err string `json:"err,omitempty"`
}

// encodeJson json解析，解析错误时返回内部错误
func encodeJSON(w http.ResponseWriter, v interface{}, log *root.Log) {

	response, err := json.Marshal(v)
	if err != nil {
		Error(w, err, http.StatusInternalServerError, log)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// NotFound 未找到记录处理.
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{}` + "\n"))
}

// JSONWithCookie 返回json，写入cookie
func JSONWithCookie(w http.ResponseWriter, v interface{}, cookie http.Cookie, log *root.Log) {
	encodeJSON(w, v, log)
	http.SetCookie(w, &cookie)
}
