package controller

import (
	"fmt"
	"net/http"
)

func TestingApi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	a := "manikanta"
	w.Write([]byte(a))
	fmt.Println("API is hitting\n", a)
}
