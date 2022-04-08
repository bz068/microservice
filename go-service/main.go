package main

import (
	"net/http"
	"log"
	"posts/utils"
	"posts/routes"
)

func main() {

	http.HandleFunc("/", utils.PathLogger(routes.RouteHandler))
	
	utils.LogServerLive()
	log.Fatal(http.ListenAndServe(":8080", nil))
}