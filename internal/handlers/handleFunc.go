package handlers

import (
	"fmt"

	"github.com/gorilla/mux"
  "net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  response := fmt.Sprintf("page %v", id)
  fmt.Fprintf(w, response)
}
