package model

import (
	"context"
	"gorm.io/gorm"
	"sync"
	"time"

	"github.com/borntodie-new/todo-list-backup/constant"
)

type User struct {
	ID        int64          `json:"id" gorm:"column:id;primarykey"`
	Username  string         `json:"username" gorm:"column:username;unique"`
	Password  string         `json:"-" gorm:"column:password"`
	Email     string         `json:"email" gorm:"column:email;unique"`
	Avatar    string         `json:"avatar" gorm:"column:avatar"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;index"`
}

func (User) TableName() string {
	return constant.UserTableName
}

type UserDao struct {
	ctx context.Context
	db  *gorm.DB
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDao(ctx context.Context, db *gorm.DB) *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{ctx: ctx, db: db}
	})
	return userDao
}

func (d *UserDao) CreateInstance(user *User) error {
	return d.db.WithContext(d.ctx).Create(user).Error
}
func (d *UserDao) RetrieveInstances(ids []int64) ([]*User, error) {
	us := make([]*User, 0)
	err := d.db.WithContext(d.ctx).Where("id in ?", ids).Find(&us).Error
	return us, err
}
func (d *UserDao) RetrieveInstance(username string) (*User, error) {
	u := new(User)
	err := d.db.WithContext(d.ctx).Where("username = ?", username).Find(&u).Error
	return u, err
}
func (d *UserDao) UpdateInstanceOfPassword(username, password string) error {
	return d.db.WithContext(d.ctx).Model(&User{}).Where("username = ?", username).Update("password", password).Error
}
