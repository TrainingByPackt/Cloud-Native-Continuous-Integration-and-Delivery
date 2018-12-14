// Package server provides functionality for handling data input and queries
package server

import (
	"context"
	"encoding/json"
	"gitlab.com/onuryilmaz/book-server/pkg/books"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/onuryilmaz/book-server/pkg/commons"
)

// REST provides functionality for HTTP REST API Server
type REST struct {
	router *httprouter.Router
	server *http.Server
	port   string
	books  books.BookDatabase

	version string
}

// NewREST creates a REST API server instance with the provided options
func NewREST(options commons.Options) *REST {
	rest := &REST{}
	rest.port = options.ServerPort
	rest.router = httprouter.New()
	rest.books = options.Books
	rest.version = options.Version
	return rest
}

// Start starts REST API server and connects handlers to the router on port
func (r *REST) Start() {

	logrus.Info("Starting REST server...")
	logrus.Infof("REST server connecting to port %v", r.port)

	r.router.GET("/ping", r.pingHandler)

	r.router.GET("/v1/init", r.initBooks)
	r.router.GET("/v1/books", r.booksHandler)

	r.server = &http.Server{Addr: ":" + r.port, Handler: r.router}
	go func() {
		err := r.server.ListenAndServe()
		if err != nil {
			logrus.Error("Error starting server: ", err)
		}
	}()
}

// Stop stops REST API gracefully
func (r *REST) Stop() {
	logrus.Warn("Stopping REST server..")
	err := r.server.Shutdown(context.TODO())
	if err != nil {
		logrus.Error("Error stopping server: ", err)
	}
}

func (r *REST) pingHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	data := map[string]string{"Status": "OK", "Version": r.version}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (r *REST) booksHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	books, err := r.books.GetBooks()

	if err != nil {
		logrus.Error("Error getting books: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(books); err != nil {
		logrus.Error("Error writing books: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (r *REST) initBooks(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	err := r.books.Initialize()

	if err != nil {
		logrus.Error("Error initialization: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}

}
