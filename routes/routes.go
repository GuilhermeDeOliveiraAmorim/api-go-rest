package routes

import (
	"log"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/go-rest-api.git/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
