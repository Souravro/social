package handler

import (
	"fmt"
	"net/http"
)

func Broadcast(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
}
