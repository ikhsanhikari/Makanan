package model




type roleModel struct {
	Id int 		`json:"id" validate:"required"`
	Role string `json:"role" validate:"required"`

}