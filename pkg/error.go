package root

import (
	"bytes"
	"fmt"
)

// Error 错误处理结构
type Error struct {
	// 错误码
	Code string

	// 错误信息，显示给终端客户用
	Message string

	// 错误详情，用于追溯
	Op  string
	Err error
}

// 通用错误码
const (
	ECONFLICT = "conflict"  // 冲突,用于当前记录与现有记录冲突
	EINTERNAL = "internal"  // 内部错误, 用于不易于对外展示的错误
	EINVALID  = "invalid"   // 未通过验证， 用于验证失败
	ENOFOUND  = "not_found" // 未找到， 用于未发现记录
)

// Error 实现错误接口
func (e *Error) Error() string {
	var buf bytes.Buffer

	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}

		buf.WriteString(e.Message)
	}

	return buf.String()
}

// ErrorCode 获取错误码
func ErrorCode(err error) string {

	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return EINTERNAL
}

// ErrorMessage 获取错误信息
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}

	return "系统内部错误，请联系系统管理员"
}
