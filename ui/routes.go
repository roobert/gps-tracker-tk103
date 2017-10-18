package ui

import (
	"net/http"

	"github.com/roobert/gps-tracker-tk103/ui/handler"
)

func SetupRoutes() {
	http.Handle("/", http.FileServer(http.Dir("ui/public/")))
	http.HandleFunc("/api", handler.API)
}
