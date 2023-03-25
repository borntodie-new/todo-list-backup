package todo

import (
	"context"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/model"
	"gorm.io/gorm"
	"sync"
)

type GetTodoFlow struct {
	// global data
	ctx context.Context
	db  *gorm.DB

	// request data
	Username string
	UserId   int64
	TodoId   int64

	// temporary data
	user *model.User
	todo *model.Todo
	// response data
	UserAndTodo *UserAndTodoDetail
}

func NewGetTodoFlow(Username string, userId, todoId int64, ctx context.Context, db *gorm.DB) *GetTodoFlow {
	return &GetTodoFlow{
		ctx:      ctx,
		db:       db,
		Username: Username,
		UserId:   userId,
		TodoId:   todoId,
	}
}

func GetTodo(username string, userId, todoId int64, ctx context.Context, db *gorm.DB) (*UserAndTodoDetail, error) {
	return NewGetTodoFlow(username, userId, todoId, ctx, db).Do()
}
func (f *GetTodoFlow) Do() (*UserAndTodoDetail, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	if err := f.packageData(); err != nil {
		return nil, err
	}
	return f.UserAndTodo, nil
}

func (f *GetTodoFlow) checkParam() error {
	if f.TodoId == 0 || f.UserId == 0 || f.Username == "" {
		return constant.ParamErr
	}
	return nil
}

func (f *GetTodoFlow) prepareData() error {
	// 获取用户信息
	wg := sync.WaitGroup{}
	wg.Add(2)
	var (
		err  error
		user *model.User
		todo *model.Todo
	)
	go func() {
		defer wg.Done()
		user, err = model.NewUserDao(f.ctx, f.db).RetrieveInstance(f.Username)
	}()
	// 获取记录信息
	go func() {
		defer wg.Done()
		todo, err = model.NewTodoDao(f.ctx, f.db).RetrieveInstance(f.UserId, f.TodoId)
	}()
	wg.Wait()
	if err != nil {
		return err
	}
	f.user = user
	f.todo = todo
	return nil
}

func (f *GetTodoFlow) packageData() error {
	f.UserAndTodo = &UserAndTodoDetail{
		DetailUser: &DetailUser{
			ID:        f.user.ID,
			Username:  f.user.Username,
			Email:     f.user.Email,
			Avatar:    f.user.Avatar,
			CreatedAt: f.user.CreatedAt,
			UpdatedAt: f.user.UpdatedAt,
		},
		Todos: []*model.Todo{f.todo},
	}
	return nil
}
