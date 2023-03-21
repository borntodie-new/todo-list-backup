package user

import (
	"context"

	"gorm.io/gorm"

	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/model"
	"github.com/borntodie-new/todo-list-backup/utils"
)

type CreateUserFlow struct {
	// global data
	ctx context.Context
	db  *gorm.DB

	// request data
	Username string
	Password string
	Email    string
	Avatar   string

	// temporarily data

	// response data
}

func CreateUser(username, password, email, avatar string ,ctx context.Context, db *gorm.DB) error  {
	return  NewCreateUserFlow(username, password, email, avatar, ctx, db).Do()
}

func NewCreateUserFlow(username, password, email, avatar string ,ctx context.Context, db *gorm.DB) *CreateUserFlow {
	return &CreateUserFlow{
		ctx:      ctx,
		db:       db,
		Username: username,
		Password: password,
		Email:    email,
		Avatar:   avatar,
	}
}

func (f *CreateUserFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.prepareData(); err != nil {
		return err
	}
	return nil
}
func (f *CreateUserFlow) checkParam() error {
	if f.Username == "" || f.Password == "" || f.Email == "" {
		return constant.ParamErr
	}
	if f.Avatar == "" {
		f.Avatar = constant.DefaultAvatarAddress
	}
	// 加密密码
	f.Password = utils.Default().GenPassword(f.Password)
	return nil
}
func (f *CreateUserFlow) prepareData() error {
	user := &model.User{
		Username: f.Username,
		Password: f.Password,
		Email:    f.Email,
		Avatar:   f.Avatar,
	}
	return model.NewUserDao(f.ctx, f.db).CreateInstance(user)
}
