package main

import (
	"fmt"
	"net/http"
	"server/config"
	"server/routes"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

//creating a map

func main() {
	config.ConnectToDB()
	routes.Routes()

	fmt.Println("SERVER RUNNING ON http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
