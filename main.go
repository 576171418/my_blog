package main

import (
	"log"
	"my_blog/routes"
	"net/http"
)

func main() {
	startWebServer("8080")
}

func startWebServer(port string) {
	r := routes.NewRouter()

	http.Handle("/", r)

	log.Printf("Starting HTTP service at " + port)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port" + port)
		log.Println("Error:" + err.Error())
	}
}
