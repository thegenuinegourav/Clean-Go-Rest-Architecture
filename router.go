package main

import (
	"github.com/Crud/Book"
	"github.com/gorilla/mux"
	"sync"
)

type router struct {}

type IRouter interface {
	InitRouter() *mux.Router
}

var(
	r 		*router
	once 	sync.Once
)

// Making Router instance as singleton
func Router() IRouter {
	if r==nil {
		once.Do(func() {
			r = &router{}
		})
	}
	return r
}

func (r *router) InitRouter() *mux.Router {
	mRouter := mux.NewRouter().StrictSlash(true)
	kernel := Kernel()
	kernel.InjectDB()
	bookController := kernel.InjectBookController()
	InitBookRoutes(mRouter, &bookController)
	return mRouter
}

func InitBookRoutes(r *mux.Router, bookController *Book.BookController) {
	r.HandleFunc("/book/{id}",bookController.GetBook).Methods("GET")
	r.HandleFunc("/book",bookController.GetBooks).Methods("GET")
	r.HandleFunc("/book",bookController.PostBook).Methods("POST")
	r.HandleFunc("/book/{id}",bookController.PutBook).Methods("PUT")
	r.HandleFunc("/book/{id}",bookController.DeleteBook).Methods("DELETE")
}
