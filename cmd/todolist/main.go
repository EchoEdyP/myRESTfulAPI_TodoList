package main

import (
	"RESTfulAPI_todos/internal/handlers"
	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/router"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

func main() {

	var cfg database.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("failed when parsing config: %v", err)
	}

	conn, err := database.ConnectDB(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	todoHandlers := handlers.NewTodoListhandlersImpl(conn)

	router := router.NewRouter(conn, &cfg, todoHandlers)

	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
