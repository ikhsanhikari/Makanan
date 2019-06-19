package main

import (
	"gopkg.in/go-playground/validator.v9"
)


type transaksiModel struct{
	Id int 					`json:"id" validate:"required"`
	IdUser string 			`json:"idUser" validate:"required"`
	KodeMakanan string 		`json:"KodeMakanan" validate:"required"`
	JumlahMakanan string 	`json:"jumlahMakanan" validate:"required"`
	CreatedAt int 			`json:"createdAt" validate:"required"`
	JumlahUang string 		`json:"jumlahUang"`
}