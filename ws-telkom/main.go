package main

import (
	"fmt"
	"log"
	"net/http"
	//"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"os"
	"github.com/gorilla/mux"
	ctr "ws-telkom/controller"
)

func main(){

	handleRequests()
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
	}
	fmt.Fprint(w, "Hello, World!")
}


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage Imran!")
    fmt.Println("Endpoint Hit: homePage")
	fmt.Println("Hello world...")
}


func handleRequests() {
	

	//===================================================================================

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/item/", ctr.GetItem).Methods("GET")
	router.HandleFunc("/", indexHandler)
	/*
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/login/{parameter}", ctr.GetLogin).Methods("GET")
	*/
	//http.HandleFunc("/login", ctr.GetLogin)

	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}

}