package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/users", users).Methods("GET")
	myRouter.HandleFunc("/getpost/{user}", getpost).Methods("GET")
	myRouter.HandleFunc("/getcomment/{user}", getcomment).Methods("GET")
	myRouter.HandleFunc("/forgetid", forgetid).Methods("GET")

	myRouter.HandleFunc("/signup", signup).Methods("POST")
	myRouter.HandleFunc("/post/{id}", postuser).Methods("POST")
	myRouter.HandleFunc("/postcomment/{user}", postcomment).Methods("POST")

	myRouter.HandleFunc("/deluser/{user}", deleteuser).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:8090", myRouter))
}

func main() {
	handleRequest()
}
