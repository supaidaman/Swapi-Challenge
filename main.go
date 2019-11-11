package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/supaidaman/go-restapi/config"
	. "github.com/supaidaman/go-restapi/dao"
	planetrouder "github.com/supaidaman/go-restapi/router"
)

var dao = PlanetsDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/planets", planetrouder.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/planets/name/{name}", planetrouder.GetPlanetsByName).Methods("GET")
	r.HandleFunc("/api/v1/planets/{id}", planetrouder.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/planets", planetrouder.Create).Methods("POST")
	r.HandleFunc("/api/v1/planets/{id}", planetrouder.Update).Methods("PUT")
	r.HandleFunc("/api/v1/planets/{id}", planetrouder.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
