package routes

import (
	"net/http"
	"social/broadcast/handler"
)

func RegisterRoutes() {
	// broadcasts a social request made by an user
	http.HandleFunc("/broadcast", handler.Broadcast)
}
