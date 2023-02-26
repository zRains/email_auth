package model

import "time"

type SignInRequired struct {
	Email    string `json:"email" bson:"Email" binding:"required,email"`
	Password string `json:"password" bson:"Password" binding:"required,min=10,max=25"`
}

type SignUpRequired struct {
	Name     string `json:"name" binding:"required,min=2,max=10"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=10,max=25"`
}

type User struct {
	Name       string    `json:"name" bson:"Name"`
	Email      string    `json:"email" bson:"Email"`
	Password   string    `json:"password" bson:"Password"`
	VerifyCode string    `json:"verifyCode,omitempty" bson:"VerifyCode,omitempty"`
	Verified   bool      `json:"verified" bson:"Verified"`
	CreateAt   time.Time `json:"createAt" bson:"CreateAt"`
	UpdateAt   time.Time `json:"updateAt" bson:"UpdateAt"`
}

type DBUserResponse struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"Name"`
	Email    string `json:"email" bson:"Email"`
	Verified bool   `json:"verified" bson:"Verified"`
}

type UserResResponse struct {
	Name     string `json:"name" bson:"Name"`
	Email    string `json:"email" bson:"Email"`
	Verified bool   `json:"verified" bson:"Verified"`
}
