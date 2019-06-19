package model

import (
	"gopkg.in/go-playground/validator.v9"
)


type roleModel struct {
	Id int `json:"id" validate:"required"`
	Role string `json:"role" validate:"required"`

}