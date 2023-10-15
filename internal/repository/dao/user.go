package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

var (
	EmailReduplicateError = errors.New("邮箱冲突")
	UserNotFindError      = gorm.ErrRecordNotFound
)

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(c context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.CTime = now
	user.UTime = now
	err := dao.db.WithContext(c).Create(&user).Error
	var me *mysql.MySQLError
	if errors.As(err, &me) {
		const duplicateErr = 1062
		if me.Number == duplicateErr {
			return EmailReduplicateError
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(c context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("email = ?", email).First(&u).Error
	return u, err
}

func (this *UserDAO) UpdateEditProfile(c context.Context, email string, name string, birthday int64, introduce string) (User, error) {
	// 更新数据库的字段
	var u User
	data := User{NickName: name, Birthday: birthday, Introduce: introduce}
	err := this.db.WithContext(c).Model(&u).Where("email = ?", email).Updates(data).Error
	return u, err
}

func (this *UserDAO) FindByUserProfile(ctx context.Context, email string) (User, error) {
	var u User
	err := this.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
}

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string

	Birthday  int64
	NickName  string
	Introduce string

	CTime int64
	UTime int64
}
