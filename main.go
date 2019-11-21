package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var userDB = map[string]int{
	"demonjoub": 18,
	"php":       80,
	"java":      20,
	"golang":    10,
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlerHomePage)
	router.HandleFunc("/signup", handlerSignup)
	router.HandleFunc("/login", handlerLogin)
	router.HandleFunc("/user/{name}", handlerUser).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func handlerSignup(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "signup.html")
}

func handlerHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// name := r.URL.Path[len("/user/"):]
	name := vars["name"]
	age := userDB[name]

	fmt.Fprintf(w, "My Name is %s, %d year old", name, age)
}
