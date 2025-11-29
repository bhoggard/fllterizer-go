package handlers

import (
	"net/http"

	"github.com/bhoggard/fllterizer-go/internal/constants"
)

func init() {
	constants.ROUTER.HandleFunc("/", helloWorldGET).Methods("GET")
}

func helloWorldGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}
