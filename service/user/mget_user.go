package user

import (
	"context"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/model"
	"gorm.io/gorm"
	"sync"
	"time"
)

type DetailUser struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailTodo struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailUserAndTodo struct {
	*DetailUser
	Todos []*model.Todo `json:"todos,omitempty"` // 空的列表返回null
}

type MGetUserFlow struct {
	// global data
	ctx context.Context
	db  *gorm.DB

	// request data
	UserIds []int64

	// temporarily data
	usersMap map[int64]*model.User
	todosMap map[int64][]*model.Todo

	// response data
	UserAndTodo []*DetailUserAndTodo
}

func MGetUser(userIds []int64, ctx context.Context, db *gorm.DB) ([]*DetailUserAndTodo, error) {
	return NewMGetUserFlow(userIds, ctx, db).Do()
}

func NewMGetUserFlow(userIds []int64, ctx context.Context, db *gorm.DB) *MGetUserFlow {
	return &MGetUserFlow{
		ctx:         ctx,
		db:          db,
		UserIds:     userIds,
		usersMap:    make(map[int64]*model.User),
		todosMap:    make(map[int64][]*model.Todo),
		UserAndTodo: make([]*DetailUserAndTodo, 0),
	}
}

func (f *MGetUserFlow) Do() ([]*DetailUserAndTodo, error) {
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

func (f *MGetUserFlow) checkParam() error {
	if len(f.UserIds) <= 0 {
		return constant.ParamErr
	}
	return nil
}

func (f *MGetUserFlow) prepareData() error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var err error
	// 查询用户记录
	go func() {
		defer wg.Done()
		users := make([]*model.User, 0)
		users, err = model.NewUserDao(f.ctx, f.db).RetrieveInstances(f.UserIds)
		for _, user := range users {
			f.usersMap[user.ID] = user
		}
	}()
	// 查询todo记录
	go func() {
		defer wg.Done()
		todos := make([]*model.Todo, 0)
		todos, err = model.NewTodoDao(f.ctx, f.db).RetrieveInstances(f.UserIds)
		for _, todo := range todos {
			if _, ok := f.todosMap[todo.UserId]; ok {
				f.todosMap[todo.UserId] = append(f.todosMap[todo.UserId], todo)
			} else {
				f.todosMap[todo.UserId] = make([]*model.Todo, 0)
				f.todosMap[todo.UserId] = append(f.todosMap[todo.UserId], todo)
			}
		}

	}()
	wg.Wait()
	return err
}

func (f *MGetUserFlow) packageData() error {
	for _, user := range f.usersMap {
		temp := &DetailUserAndTodo{
			DetailUser: &DetailUser{
				Id:        user.ID,
				Username:  user.Username,
				Email:     user.Email,
				Avatar:    user.Avatar,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			Todos: make([]*model.Todo, 0),
		}
		if todos, ok := f.todosMap[user.ID]; ok {
			temp.Todos = todos
		}
		f.UserAndTodo = append(f.UserAndTodo, temp)
	}
	return nil
}
