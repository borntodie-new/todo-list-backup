package todo

import (
	"context"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/model"
	"gorm.io/gorm"
)

type UpdateTodoFlow struct {
	// global
	ctx context.Context
	db  *gorm.DB

	// request data
	UserId  int64
	TodoId  int64
	Content string

	// response data
}

func NewUpdateTodoFlow(userId, todoId int64, content string, ctx context.Context, db *gorm.DB) *UpdateTodoFlow {
	return &UpdateTodoFlow{
		ctx:     ctx,
		db:      db,
		UserId:  userId,
		TodoId:  todoId,
		Content: content,
	}
}

func UpdateTodo(userId, todoId int64, content string, ctx context.Context, db *gorm.DB) error {
	return NewUpdateTodoFlow(userId, todoId, content, ctx, db).Do()
}

func (f *UpdateTodoFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.prepareData(); err != nil {
		return err
	}
	return nil
}

func (f *UpdateTodoFlow) checkParam() error {
	if f.UserId == 0 || f.TodoId == 0 || f.Content == "" {
		return constant.ParamErr
	}
	return nil
}
func (f *UpdateTodoFlow) prepareData() error {
	err := model.NewTodoDao(f.ctx, f.db).UpdateInstance(f.UserId, f.TodoId, f.Content)
	return err
}
