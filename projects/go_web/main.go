package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yush2411/bazelisk-task/projects/go_hello_world"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	r := Println("Received request")
	w.Write([]byte(go_hello_world.HelloWorld()))
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)

	log.Println("Going to listen on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}