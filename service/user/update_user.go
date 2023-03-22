package user

import (
	"context"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/model"
	"github.com/borntodie-new/todo-list-backup/utils"
	"gorm.io/gorm"
)

type UpdateUserFlow struct {
	// global data
	ctx context.Context
	db  *gorm.DB

	// request data
	Username    string
	NewPassword string

	// temporarily data
	// response data
}

func UpdateUser(username, newPassword string, ctx context.Context, db *gorm.DB) error {
	return NewUpdateUserFlow(username, newPassword, ctx, db).Do()
}

func NewUpdateUserFlow(username, newPassword string, ctx context.Context, db *gorm.DB) *UpdateUserFlow {
	return &UpdateUserFlow{
		ctx:         ctx,
		db:          db,
		Username:    username,
		NewPassword: newPassword,
	}
}

func (f *UpdateUserFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.prepareData(); err != nil {
		return err
	}
	return nil
}
func (f *UpdateUserFlow) checkParam() error {
	if f.Username == "" || f.NewPassword == "" {
		return constant.ParamErr
	}
	f.NewPassword = utils.Default().GenPassword(f.NewPassword)
	return nil
}
func (f *UpdateUserFlow) prepareData() error {
	if err := model.NewUserDao(f.ctx, f.db).UpdateInstanceOfPassword(f.Username, f.NewPassword); err != nil {
		return err
	}
	return nil
}
