package model

import (
	"context"
	"gorm.io/gorm"
	"sync"
	"time"

	"github.com/borntodie-new/todo-list-backup/constant"
)

type Todo struct {
	ID        int64          `json:"id" gorm:"column:id;primarykey"`
	UserId    int64          `json:"user_id" gorm:"column:user_id"`
	Context   string         `json:"username" gorm:"column:context"`
	Completed bool           `json:"completed" gorm:"column:completed;default:false"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:updated_at;index"`
}

func (*Todo) TableName() string {
	return constant.TodoTableName
}

type TodoDao struct {
	ctx context.Context
	db  *gorm.DB
}

var (
	todoDao  *TodoDao
	todoOnce sync.Once
)

func NewTodoDao(ctx context.Context, db *gorm.DB) *TodoDao {
	todoOnce.Do(func() {
		todoDao = &TodoDao{ctx: ctx, db: db}
	})
	return todoDao
}

func (d *TodoDao) CreateInstance(todo *TodoDao) error {
	return d.db.WithContext(d.ctx).Create(&todo).Error
}
func (d *TodoDao) DeleteInstance(id int64) error {
	return d.db.WithContext(d.ctx).Where("id = ?", id).Update("completed", true).Error
}
func (d *TodoDao) RetrieveInstances(userId int64, offset, limit int) ([]*Todo, error) {
	ts := make([]*Todo, 0)
	err := d.db.WithContext(d.ctx).Where("user_id = ?", userId).Limit(limit).Offset(offset).Order("age create_at").Find(&ts).Error
	return ts, err
}
func (d *TodoDao) UpdateInstance(id int64, content string) error {
	return d.db.WithContext(d.ctx).Where("id = ?", id).Update("context", content).Error
}