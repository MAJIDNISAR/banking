package app

import (
	"log"
	"net/http"
)

func Start() {
	//functional programing syntax
	// http.HandleFunc("/greet", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Greet Request hit")
	// 	fmt.Fprint(rw, "Hello World")
	// })
	// we can define our own request multiplexer or router
	router := http.NewServeMux()

	// defining the routes using router and not http
	// http.HandleFunc("/greet", greet) it was done previously like this
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)

	// starting the server and listeing on the port and log.fatal in case of error starting the server will be logged
	// log.Fatal(http.ListenAndServe("localhost:8000", nil)) default multiplexer or router
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
