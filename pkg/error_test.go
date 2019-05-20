package root_test

import (
	"errors"
	"fmt"
	"testing"
	"b2yun/pkg"
)

func TestError_ErrorCode(t *testing.T) {

	// 测试指定错误码
	err := root.Error{
		Code: root.ENOFOUND,
	}

	if code := root.ErrorCode(&err); code != root.ENOFOUND {
		t.Errorf("错误码检测失败，期待%s, 实际%s", root.ENOFOUND, code)
		return
	}

	// 测试嵌套错误码
	errWrapper := root.Error{
		Err: &err,
	}

	if code := root.ErrorCode(&errWrapper); code != root.ENOFOUND {
		t.Errorf("错误码检测失败，期待[%s], 实际[%s]", root.ENOFOUND, code)
		return
	}

	// 测试无Code
	if code := root.ErrorCode(nil); code != "" {
		t.Errorf("nil值错误码检测失败,期待[%s], 实际[%s]", "", code)
		return
	}

	// 测试非root.Error类型
	errStd := errors.New("标准错误类型")
	if code := root.ErrorCode(errStd); code != root.EINTERNAL {
		t.Errorf("标准错误检测失败m期待[%s], 实际[%s]", root.EINTERNAL, code)
		return
	}

}

func TestErro_ErrorMessage(t *testing.T) {

	// 检测错误信息
	err := root.Error{
		Message: "err",
	}

	if msg := root.ErrorMessage(&err); msg != "err" {
		t.Errorf("错误信息检测错误，期待[%s], 实际[%s]", "err", msg)
		return
	}

	// 检测嵌套错误
	errWrapper := root.Error{
		Err: &err,
	}

	if msg := root.ErrorMessage(&errWrapper); msg != "err" {
		t.Errorf("嵌套错误信息检测错误，期待[%s], 实际[%s]", "err", msg)
		return
	}

	// 检测空错误
	if msg := root.ErrorMessage(nil); msg != "" {
		t.Errorf("空错误检测错误，期待[%s], 实际[%s]", "", msg)
	}

	// 检测标准错误
	const stdMessage = "系统内部错误，请联系系统管理员"
	errStd := errors.New("std")
	if msg := root.ErrorMessage(errStd); msg != stdMessage {
		t.Errorf("标准错误嵌套错误，期待[%s], 实际[%s]", stdMessage, msg)
	}

}

func TestError_Error(t *testing.T) {

	const op = "op"
	const msg = "msg"
	const code = "code"
	const Err = "err"

	err := root.Error{
		Op:      op,
		Message: msg,
		Code:    code,
		Err:     errors.New(Err),
	}

	if errMsg := err.Error(); errMsg != fmt.Sprintf("%s: %s", op, Err) {
		t.Errorf("error解析错误，期待[%s], 实际[%s]", fmt.Sprintf("%s: %s", op, Err), errMsg)
	}

	errMsgCode := root.Error{
		Op:      op,
		Code:    code,
		Message: msg,
	}

	if errMsg := errMsgCode.Error(); errMsg != fmt.Sprintf("%s: <%s> %s", op, code, msg) {
		t.Errorf("error解析错误，期待[%s: <%s> %s], 实际[%s]", op, code, msg, errMsg)
	}

	errMsgOp := root.Error{
		Op: op,
	}
	if errMsgOp := errMsgOp.Error(); errMsgOp != "op: " {
		t.Errorf("errorop解析错误")
	}

}
