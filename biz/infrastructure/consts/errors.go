package consts

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Errno struct {
	err  error
	code codes.Code
}

// GRPCStatus 实现 GRPCStatus 方法
func (en *Errno) GRPCStatus() *status.Status {
	return status.New(en.code, en.err.Error())
}

// 实现 Error 方法
func (en *Errno) Error() string {
	return en.err.Error()
}

// NewErrno 创建自定义错误
func NewErrno(code codes.Code, err error) *Errno {
	return &Errno{
		err:  err,
		code: code,
	}
}

// 定义常量错误
var (
	ErrNotAuthentication = NewErrno(codes.Code(1000), errors.New("not authentication"))
	ErrForbidden         = NewErrno(codes.PermissionDenied, errors.New("forbidden"))
	ErrSignUp            = NewErrno(codes.Code(1001), errors.New("注册失败，请重试"))
	ErrSignIn            = NewErrno(codes.Code(1002), errors.New("登录失败，请先注册或重试"))
	ErrRepeatedSignUp    = NewErrno(codes.Code(1004), errors.New("该手机号已注册"))
	ErrNotSignUp         = NewErrno(codes.Code(1006), errors.New("请确认手机号已注册"))
	ErrSend              = NewErrno(codes.Code(1007), errors.New("发送验证码失败，请重试"))
	ErrVerifyCode        = NewErrno(codes.Code(1008), errors.New("验证码错误"))
	ErrDefaultGroup      = NewErrno(codes.Code(1009), errors.New("默认组创建失败"))
	ErrCall              = NewErrno(codes.Code(1010), errors.New("模型调用失败"))
)

// 数据库相关错误
var (
	ErrNotFound        = NewErrno(codes.NotFound, errors.New("not found"))
	ErrInvalidObjectId = NewErrno(codes.InvalidArgument, errors.New("无效的id "))
	ErrUpdate          = NewErrno(codes.Code(2001), errors.New("更新失败"))
	ErrInsert          = NewErrno(codes.Code(2002), errors.New("创建失败"))
)
