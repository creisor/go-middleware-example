package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	n := negroni.Classic() // Includes some default middlewares
	n.Use(negroni.HandlerFunc(verifyMiddleware))
	n.UseHandler(router)

	http.ListenAndServe(":3000", n)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func verifyMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	auth := r.Header.Get("auth")
	log.Printf("auth: %s", auth)
	if auth == "notahacker" {
		next(w, r)
		return
	}

	log.Printf("Not authorized")
	http.Error(w, "Not authorized", http.StatusUnauthorized)
}
