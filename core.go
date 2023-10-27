package main

import (
	"arrieup/collocom/database"
	"arrieup/collocom/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Main start")
	server.SetupRoutes()
	database.DBsetup()
	fmt.Println(database.ReadUserByUsername("arrieup"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
