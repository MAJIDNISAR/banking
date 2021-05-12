package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Data layer
type Customer struct {
	Name    string `json:"full_name" xml:"Name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

// Request Handlers
func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Greet Request hit")
	fmt.Fprint(rw, "Hello World")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllCustomers Request hit")
	customers := []Customer{
		{Name: "Majid Nisar", City: "Srinagar", Zipcode: "190003"},
		{Name: "Hamida Majid", City: "Srinagar", Zipcode: "190003"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
