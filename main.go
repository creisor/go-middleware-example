package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	key := "notahacker"

	router := mux.NewRouter()
	authRoutes := mux.NewRouter()

	common := negroni.Classic()

	router.PathPrefix("/authorized").Handler(common.With(
		NewVerifyMiddleware(key),
		NewGenericHandler("You are authorized"),
		negroni.Wrap(authRoutes),
	))

	// order matters (/authorized must come before /, or / will match the /authorized route)
	router.PathPrefix("/").Handler(common.With(
		NewGenericHandler("Hello from home"),
	))

	n := negroni.New()
	n.UseHandler(router)

	http.ListenAndServe(":3000", n)
}

type GenericHandler struct {
	text string
}

func NewGenericHandler(text string) *GenericHandler {
	return &GenericHandler{text}
}

func (h *GenericHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Fprintf(w, h.text)
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
