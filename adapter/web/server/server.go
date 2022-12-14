package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ffelipelimao/ports-adapters-architecture/adapter/web/handler"
	"github.com/ffelipelimao/ports-adapters-architecture/application"
	"github.com/gorilla/mux"

	"github.com/urfave/negroni"
)

type WebServer struct {
	Service application.IProductService
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductsHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
