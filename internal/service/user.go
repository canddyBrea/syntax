package service

import (
	"context"
	"errors"

	"syntax/internal/model"
	"syntax/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

var (
	EmailReduplicateError      = repository.EmailReduplicateError
	ErrorInvalidUserOrPassword = errors.New("用户不存在或密码不对")
)

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Signup 注册逻辑管控 (register logic controller area)
func (us *UserService) Signup(c context.Context, u *model.User) error {
	enCryption, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(enCryption)
	return us.repo.Create(c, *u)
}

func (us *UserService) Login(c context.Context, email string, password string) (model.User, error) {
	u, err := us.repo.FindByEmail(c, email)
	if errors.Is(err, repository.UserNotFindError) {
		return model.User{}, ErrorInvalidUserOrPassword
	}
	if err != nil {
		return model.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return model.User{}, ErrorInvalidUserOrPassword
	}
	return u, nil
}

func (this *UserService) Edit(c context.Context, email string, name string, birthday int64, introduce string) (model.SendUserProfile, error) {
	user, err := this.repo.UpdateEditProfile(c, email, name, birthday, introduce)
	if err != nil {
		return model.SendUserProfile{}, err
	}
	return user, nil
}

func (this *UserService) Profile(c context.Context, email string) (model.SendUserProfile, error) {
	user, err := this.repo.FindByUserProfile(c, email)
	if err != nil {
		return model.SendUserProfile{}, err
	}
	return user, nil
}
