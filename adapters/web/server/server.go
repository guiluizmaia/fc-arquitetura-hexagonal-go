package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/guiluizmaia/fc2-arquitetura-hexagonal-go/application"
	"github.com/urfave/negroni"
)

type Webserver struct {
	Service application.IProductService
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Server() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           nil,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
