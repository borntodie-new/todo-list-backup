package todo

import (
	"context"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/model"
	"gorm.io/gorm"
	"sync"
	"time"
)

type DetailUser struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserAndTodoDetail struct {
	*DetailUser
	Todos []*model.Todo `json:"todos,omitempty"`
}

type MGetTodoFlow struct {
	// global data
	ctx context.Context
	db  *gorm.DB

	// request data
	UserIds []int64

	// temporary data
	usersMapping map[int64]*model.User
	todosMapping map[int64][]*model.Todo

	// response data
	UserAndTodo []*UserAndTodoDetail
}

func NewMGetTodoFlow(userIds []int64, ctx context.Context, db *gorm.DB) *MGetTodoFlow {
	return &MGetTodoFlow{
		ctx:          ctx,
		db:           db,
		UserIds:      userIds,
		usersMapping: make(map[int64]*model.User),
		todosMapping: make(map[int64][]*model.Todo),
		UserAndTodo:  make([]*UserAndTodoDetail, 0),
	}
}
func MGetTodo(UserIds []int64, ctx context.Context, db *gorm.DB) ([]*UserAndTodoDetail, error) {
	return NewMGetTodoFlow(UserIds, ctx, db).Do()
}

func (f *MGetTodoFlow) Do() ([]*UserAndTodoDetail, error) {
	if err := f.checkData(); err != nil {
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
func (f *MGetTodoFlow) checkData() error {
	if len(f.UserIds) == 0 {
		return constant.ParamErr
	}
	return nil
}

func (f *MGetTodoFlow) prepareData() error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var err error
	// 1. 获取用户列表
	go func() {
		defer wg.Done()
		var users []*model.User
		users, err = model.NewUserDao(f.ctx, f.db).RetrieveInstances(f.UserIds)
		for _, user := range users {
			f.usersMapping[user.ID] = user
		}
	}()

	// 2. 获取todo列表
	go func() {
		defer wg.Done()
		var todos []*model.Todo
		todos, err = model.NewTodoDao(f.ctx, f.db).RetrieveInstances(f.UserIds)
		for _, todo := range todos {
			if _, ok := f.todosMapping[todo.UserId]; ok {
				f.todosMapping[todo.UserId] = append(f.todosMapping[todo.UserId], todo)
			} else {
				f.todosMapping[todo.UserId] = make([]*model.Todo, 0)
				f.todosMapping[todo.UserId] = append(f.todosMapping[todo.UserId], todo)
			}
		}
	}()

	wg.Wait()
	return err
}

func (f *MGetTodoFlow) packageData() error {
	userAndTodoDetails := make([]*UserAndTodoDetail, 0)
	for _, user := range f.usersMapping {
		temp := &UserAndTodoDetail{
			DetailUser: &DetailUser{
				ID:        user.ID,
				Username:  user.Username,
				Email:     user.Email,
				Avatar:    user.Avatar,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			Todos: make([]*model.Todo, 0),
		}
		if todos, ok := f.todosMapping[user.ID]; ok {
			temp.Todos = todos
		}
		userAndTodoDetails = append(userAndTodoDetails, temp)
	}
	f.UserAndTodo = userAndTodoDetails
	return nil
}
