package main

import (
	"net/http"

	. "github.com/roobert/golang-db"
	_ "github.com/roobert/gps-tracker-tk103/common"
	. "github.com/roobert/gps-tracker-tk103/ui"
)

func main() {
	defer DB.Close()
	SetupRoutes()
	http.ListenAndServe(":9601", nil)
}
