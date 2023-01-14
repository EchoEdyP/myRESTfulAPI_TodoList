package main

import (
	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/router"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	// Use a MySQLConn struct as the implementation of the DBConn interface when calling InsertTodos
	conn := &database.MySQLConn{}
	// Create a new router with the appropriate routes
	router := router.NewRouter(conn)

	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
