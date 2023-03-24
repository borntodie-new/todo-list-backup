package handler

import (
	"github.com/borntodie-new/todo-list-backup/constant"
)

type BaseResp struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func RespSuccess() *BaseResp {
	return &BaseResp{
		StatusCode: constant.SuccessCode,
		Message:    constant.Success,
	}
}

func RespSuccessWithData(data interface{}) *BaseResp {
	return &BaseResp{
		StatusCode: constant.SuccessCode,
		Message:    constant.Success,
		Data:       data,
	}
}

func RespFailed(err error) *BaseResp {
	code, message := errorConvertToCodeAndMsg(err)
	return &BaseResp{
		StatusCode: code,
		Message:    message,
	}
}

func RespFailedWithData(err error, data interface{}) *BaseResp {
	code, message := errorConvertToCodeAndMsg(err)
	return &BaseResp{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}
}

func errorConvertToCodeAndMsg(err error) (int, string) {
	msg := err.Error()
	switch err {
	case constant.ParamErr:
		return constant.ParamErrCode, msg
	case constant.CodeExpiresErr:
		return constant.CodeExpiresErrCode, msg
	case constant.CodeIncorrectErr:
		return constant.CodeIncorrectErrCode, msg
	case constant.UserPasswordErr:
		return constant.UserPasswordErrCode, msg
	case constant.UserAlreadyExistErr:
		return constant.UserAlreadyExistErrCode, msg
	case constant.AuthorizationFailedErr:
		return constant.AuthorizationFailedErrCode, msg
	case constant.TokenExpiredErr:
		return constant.TokenExpiredErrCode, msg
	case constant.TokenInvalidErr:
		return constant.TokenInvalidErrCode, msg
	default:
		return constant.ServiceErrCode, msg
	}
}
