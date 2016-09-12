package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/hello", HelloHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

var store = sessions.NewCookieStore([]byte("very-secret"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession");
	
	session.Values["Name"] = "Tam"
	
	fmt.Fprint(w, "Index")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}
