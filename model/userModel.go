package model

import (
	"gopkg.in/go-playground/validator.v9"
)


type UserModel struct {
	Id int 			`json:"id" validate:"required"`
	NamaUser string `json:"namaUser" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	IdRole int 		`json:"idRole" validate:"required"`
	Phone string 	`json:"phone"`
	Email string 	`json:"email"`

}