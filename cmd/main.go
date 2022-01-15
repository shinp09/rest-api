package main

import (
	"github.com/koga456/sample-api/controller"
	"github.com/koga456/sample-api/model/repository"
	"net/http"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe()
}
