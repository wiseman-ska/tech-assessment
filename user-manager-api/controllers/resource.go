package controllers

import "models"

type UserResource struct {
	Data models.User `json:"data"`
}

type LoginResource struct {
	Data LoginModel `json:"data"`
}

type AuthUserResource struct {
	Data AuthUserModel `json:"data"`
}

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserModel struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}
