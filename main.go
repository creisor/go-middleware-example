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
	authRouter := mux.NewRouter().PathPrefix("/authorized").Subrouter().StrictSlash(true)

	router.HandleFunc("/", homeHandler)

	router.PathPrefix("/authorized").Handler(negroni.New(
		NewVerifyMiddleware("notahacker"),
		&AuthHandler{},
		negroni.Wrap(authRouter),
	))

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":3000", n)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from home")
}

type AuthHandler struct{}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Fprintf(w, "You are authorized")
}

type VerifyMiddleware struct {
	key string
}

func NewVerifyMiddleware(key string) *VerifyMiddleware {
	return &VerifyMiddleware{key}
}

func (v *VerifyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	auth := r.Header.Get("auth")
	log.Printf("auth: %s", auth)
	if auth == v.key {
		next(w, r)
		return
	}

	log.Printf("Not authorized")
	http.Error(w, "Not authorized", http.StatusUnauthorized)
}
