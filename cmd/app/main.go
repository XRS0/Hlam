package main

import (
	"fmt"
	"net/http"
  "Project/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../static/"))))
  router := mux.NewRouter()
  http.Handle("/", router)
  router.HandleFunc("/mainPage/{id:[0-9]+}", handlers.MainPageHandler)
  router.HandleFunc("/dataPage", handlers.DataParseHandler)
  fmt.Println("Server is listening...")
  http.ListenAndServe(":8080", nil)

}
