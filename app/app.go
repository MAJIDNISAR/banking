package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// defining the routes using router and not http
	// http.HandleFunc("/greet", greet) it was done previously like this
	router.HandleFunc("/greet", greet)                                                        // by default method is get
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)                  // using methods explicitly
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet) // regular expression for numbers only
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)                  // using methods explicitly
	// starting the server and listeing on the port and log.fatal in case of error starting the server will be logged
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
