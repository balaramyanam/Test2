package main

import (
	"inter/database"
	"inter/routes"
)

// C:\Users\jvt\go\pkg\mod\src\routes

//	"model"

func main() {

	database.Connectme()
	routes.Register()

}
