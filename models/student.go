package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validade:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg" validade:"len=9, regexp=^[0-9]*$"`
	gorm.Model
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
