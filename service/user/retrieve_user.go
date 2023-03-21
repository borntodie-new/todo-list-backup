package user

import (
	"context"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/model"
	"github.com/borntodie-new/todo-list-backup/utils"
	"gorm.io/gorm"
)

type RetrieveUserFlow struct {
	// global
	ctx context.Context
	db  *gorm.DB

	// request data
	Username string
	Password string

	// response data
	User *model.User
}

func NewRetrieveUserFlow(username, password string, ctx context.Context, db *gorm.DB) *RetrieveUserFlow {
	return &RetrieveUserFlow{
		ctx:      ctx,
		db:       db,
		Username: username,
		Password: password,
	}
}

func RetrieveUser(username, password string, ctx context.Context, db *gorm.DB) (*model.User, error) {
	return NewRetrieveUserFlow(username, password, ctx, db).Do()
}

func (f *RetrieveUserFlow) Do() (*model.User, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	return f.User, nil
}

func (f *RetrieveUserFlow) checkParam() error {
	if f.Username == "" || f.Password == "" {
		return constant.ParamErr
	}
	return nil
}

func (f *RetrieveUserFlow) prepareData() error {
	instance, err := model.NewUserDao(f.ctx, f.db).RetrieveInstance(f.Username)
	if err != nil {
		return err
	}
	verify := utils.Default().Verify(f.Password, instance.Password)
	if !verify {
		return constant.UserPasswordErr
	}
	f.User = instance
	return nil
}
