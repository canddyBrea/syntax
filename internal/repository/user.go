package repository

import (
	"context"
	"time"

	"syntax/internal/model"
	"syntax/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

var (
	EmailReduplicateError = dao.EmailReduplicateError
	UserNotFindError      = dao.UserNotFindError
)

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (repo *UserRepository) Create(c context.Context, u model.User) error {
	return repo.dao.Insert(c, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (repo *UserRepository) FindByEmail(c context.Context, email string) (model.User, error) {
	user, err := repo.dao.FindByEmail(c, email)
	if err != nil {
		return model.User{}, err
	}
	return repo.toModel(user), nil
}

func (repo *UserRepository) toModel(u dao.User) model.User {
	return model.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (this *UserRepository) UpdateEditProfile(ctx context.Context, email string, name string, birthday int64, introduce string) (model.SendUserProfile, error) {
	user, err := this.dao.UpdateEditProfile(ctx, email, name, birthday, introduce)
	if err != nil {
		return model.SendUserProfile{}, err
	}
	return this.toProfileModel(user), nil
}

func (this *UserRepository) toProfileModel(u dao.User) model.SendUserProfile {
	return model.SendUserProfile{
		Email:     u.Email,
		Birthday:  time.UnixMilli(u.Birthday).Format("2006-01-02"),
		NickName:  u.NickName,
		Introduce: u.Introduce,
	}
}

func (this *UserRepository) FindByUserProfile(ctx context.Context, email string) (model.SendUserProfile, error) {
	user, err := this.dao.FindByUserProfile(ctx, email)
	if err != nil {
		return model.SendUserProfile{}, err
	}
	return this.toProfileModel(user), nil
}
