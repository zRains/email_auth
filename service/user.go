package service

import (
	"context"
	"email_auth/data"
	"email_auth/model"
	"email_auth/util"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserService interface {
	SignUp(user *model.User) (*model.DBUserResponse, error)
	SignIn(user *model.SignInRequired) (*model.DBUserResponse, error)
	//FindUserByEmail(email string) (*model.DBUserResponse, error)
}

type UserServiceImpl struct {
	ctx context.Context
}

func CreateUserService(ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		ctx,
	}
}

func (this UserServiceImpl) SignUp(user *model.User) (*model.DBUserResponse, error) {
	now := time.Now()
	user.CreateAt = now
	user.UpdateAt = now
	user.Email = "2296342883@qq.com"
	user.Verified = false

	// Encrypt user password
	encryptedPassword, err := util.EncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = encryptedPassword

	//	Add a new user
	result, err := data.AddUser(this.ctx, user)
	if err != nil {
		if mongoError, ok := err.(mongo.WriteException); ok && mongoError.WriteErrors[0].Code == 11000 {
			return nil, errors.New("duplicate email: " + user.Email)
		}

		return nil, err
	}

	newUser, err := data.GetUserById(this.ctx, result.InsertedID)
	if err != nil {
		return nil, err
	}

	var userRes model.DBUserResponse

	newUser.Decode(&userRes)

	return &userRes, nil
}

func (this UserServiceImpl) SignIn(user *model.SignInRequired) (*model.DBUserResponse, error) {
	// Compare password
	dbResult, _ := data.GetUserByEmail(this.ctx, user.Email)

	var dbUser model.User

	dbResult.Decode(&dbUser)

	if err := util.VerifyPassword(user.Password, dbUser.Password); err != nil {
		return nil, err
	}

	var userRes model.DBUserResponse

	dbResult.Decode(&userRes)

	return &userRes, nil
}
