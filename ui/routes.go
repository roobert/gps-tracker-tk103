package ui

import (
	"net/http"
)

func SetupRoutes() {
	//http.HandleFunc("/", handler.Home)
	//http.HandleFunc("/api", handler.API)
	//http.Handle("/index.html", http.FileServer(http.Dir("./ui/html/index.html")))
	//http.Handle("/", http.StripPrefix("html", http.FileServer(http.Dir("./html"))))
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../html"))))
	http.Handle("/", http.FileServer(http.Dir("ui/public/")))
}
