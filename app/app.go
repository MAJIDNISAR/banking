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

	// defining the routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting the server and listeing on the port and log.fatal in case of error starting the server will be logged
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
