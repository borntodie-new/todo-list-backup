package todo

import (
	"context"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/model"
	"gorm.io/gorm"
)

type CreateTodoFlow struct {
	// global data
	ctx context.Context
	db  *gorm.DB

	// request data
	UserId  int64
	Content string

	// temporal data

	// response data
}

func NewCreateTodoFlow(userId int64, content string, ctx context.Context, db *gorm.DB) *CreateTodoFlow {
	return &CreateTodoFlow{
		ctx:     ctx,
		db:      db,
		UserId:  userId,
		Content: content,
	}
}

func CreateTodo(userId int64, content string, ctx context.Context, db *gorm.DB) error {
	return NewCreateTodoFlow(userId, content, ctx, db).Do()
}

func (f *CreateTodoFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.prepareData(); err != nil {
		return err
	}
	return nil
}

func (f *CreateTodoFlow) checkParam() error {
	if f.UserId <= 0 || f.Content == "" {
		return constant.ParamErr
	}
	return nil
}
func (f *CreateTodoFlow) prepareData() error {
	todo := &model.Todo{
		UserId:  f.UserId,
		Content: f.Content,
	}
	err := model.NewTodoDao(f.ctx, f.db).CreateInstance(todo)
	return err
}
