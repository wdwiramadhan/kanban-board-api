package repository

import (
	"context"

	"github.com/wdwiramadhan/kanban-board-api/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &UserRepository{Conn}
}

func(u *UserRepository) StoreUser(ctx context.Context, user *domain.User) (userId int64, err error) {
	err = u.Conn.Create(user).Error
	if err != nil {
		return 
	}
	userId = user.ID
	return 
}

func(u *UserRepository) GetUserByID(ctx context.Context, id int64) (domain.User, error) {
	var user domain.User
	err := u.Conn.First(&user,"id = ?",id).Error
	if err != nil {
		return  domain.User{},err
	}
	return user, nil
}

func(u *UserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := u.Conn.First(&user, "email = ?", email).Error
	if err != nil {
		return  domain.User{}, err
	}
	return user, nil
}

func(u *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (error){
	err := u.Conn.Model(user).Updates(user).Error
	return err
}

func(u *UserRepository) DeleteUser(ctx context.Context, id int64) (error){
	user := domain.User{ID: id}
	err := u.Conn.Delete(&user).Error
	return err
}