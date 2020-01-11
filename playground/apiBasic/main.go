package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mingsterism/go-distributed-sys/playground/apiBasic/protos"
)

// type Person struct {
// 	Name string
// 	Age  int
// }

func personCreate(w http.ResponseWriter, r *http.Request) {
	// // Declare a new Person struct.
	var p apiBasic.protos.Person
	// fmt.Println("========", r.Body)

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	fmt.Println(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...
	fmt.Fprintf(w, "Person: %+v", p)

}

func main() {
	// h1 := protos.Human{
	// 	Name: "Lukas",
	// }
	// p1 := protos.Person{
	// 	Id:    1234,
	// 	Name:  "John Doe",
	// 	Email: "jdoe@example.com",
	// 	Phones: []*protos.Person_PhoneNumber{
	// 		{Number: "555-4321", Type: protos.Person_HOME},
	// 	},
	// }
	// fmt.Println(h1)
	// fmt.Println(p1)

	mux := http.NewServeMux()
	mux.HandleFunc("/person/create", personCreate)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
