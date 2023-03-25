package todo

import (
	"context"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/model"
	"gorm.io/gorm"
)

type DeleteTodoFlow struct {
	// global data
	ctx context.Context
	db  *gorm.DB
	// request data
	UserId int64
	TodoId int64

	// temporal data

	// response data

}

func NewDeleteTodoFlow(userId, todoId int64, ctx context.Context, db *gorm.DB) *DeleteTodoFlow {
	return &DeleteTodoFlow{
		ctx:    ctx,
		db:     db,
		UserId: userId,
		TodoId: todoId,
	}
}

func DeleteTodo(userId, todoId int64, ctx context.Context, db *gorm.DB) error {
	return NewDeleteTodoFlow(userId, todoId, ctx, db).Do()
}

func (f *DeleteTodoFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.prepareData(); err != nil {
		return err
	}
	return nil
}

func (f *DeleteTodoFlow) checkParam() error {
	if f.UserId == 0 || f.TodoId == 0 {
		return constant.ParamErr
	}
	return nil
}

func (f *DeleteTodoFlow) prepareData() error {
	err := model.NewTodoDao(f.ctx, f.db).DeleteInstance(f.UserId, f.TodoId)
	return err
}

