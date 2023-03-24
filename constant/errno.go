package constant

import "errors"

var (
	Success                = "Request success"
	ServiceErr             = errors.New("Service is unable to start successfully")
	ParamErr               = errors.New("Wrong Parameter has been given")
	CodeExpiresErr         = errors.New("Code is expired")
	CodeIncorrectErr       = errors.New("Code is incorrect")
	UserPasswordErr        = errors.New("User password wrong")
	UserAlreadyExistErr    = errors.New("User already exists")
	AuthorizationFailedErr = errors.New("Authorization failed")
	TokenExpiredErr        = errors.New("Token is expired")
	TokenInvalidErr        = errors.New("Token is invalid")
)

var (
	SuccessCode                = 1000
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	CodeExpiresErrCode         = 10003
	CodeIncorrectErrCode       = 10004
	UserPasswordErrCode        = 10005
	UserAlreadyExistErrCode    = 10006
	AuthorizationFailedErrCode = 10007
	TokenExpiredErrCode        = 10008
	TokenInvalidErrCode        = 10009
)
