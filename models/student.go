package models

import "gorm.io/gorm"

type Student struct {
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
	gorm.Model
}

var Students []Student
