package dto

import(
  "github.com/vietstars/postgres-api/models"
)

type UserNew struct {
  UserName string `json:"username" validate:"required"`
  Email string `json:"email" validate:"required,email"`
  Password string `json:"password" validate:"required"`
}

type UserEdit struct {
  UserName string `json:"username"`
  Email string `json:"email" validate:"email"`
  Password string `json:"password"`
  Version int `json:"version" validate:"required"`
}

type UserSignIn struct {
  Email string `json:"email" validate:"required,email"`
  Password string `json:"password"`
}

type Auth struct {
  Info *models.UserEntity `json:"info"`
  Token string `json:"token"`
}
