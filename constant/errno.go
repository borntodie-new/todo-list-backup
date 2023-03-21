package constant

import "errors"

var (
	ServiceErr             = errors.New("Service is unable to start successfully")
	ParamErr               = errors.New("Wrong Parameter has been given")
	UserPasswordErr        = errors.New("User password wrong")
	UserAlreadyExistErr    = errors.New("User already exists")
	AuthorizationFailedErr = errors.New("Authorization failed")
)
