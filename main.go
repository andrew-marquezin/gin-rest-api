package main

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"gin-rest-api/routes"
)

func main() {
	database.ConnectWithDatabase()
	models.Students = []models.Student{
		{Name: "Andrew Marinho", CPF: "00000000000", RG: "010000000"},
		{Name: "Mell", CPF: "00000000001", RG: "020000000"},
	}
	routes.HandleRequests()
}
