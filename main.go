package main

import (
	"golang_tugas_3/db"
	"golang_tugas_3/server"
	"golang_tugas_3/server/controllers"
)

func main() {
	db := db.ConnectGorm()
	peopleController := controllers.NewPersonController(db)
	router := server.NewRouter(peopleController)
	router.Start(":4000")
}
