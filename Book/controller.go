package Book

import (
	"database/sql"
	"encoding/json"
	common "github.com/Crud/Common"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type BookController struct{
	IBookService
	common.ICommonController
}

func (bookController *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		bookController.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	book, err := bookController.GetBookService(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			bookController.RespondWithError(w, http.StatusNotFound, "Book not found")
		default:
			bookController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	bookController.RespondWithJSON(w, http.StatusOK, book)
}

func (bookController *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bookController.GetBooksService()
	if err != nil {
		bookController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	bookController.RespondWithJSON(w, http.StatusOK, books)
}

func (bookController *BookController) PostBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		bookController.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	res, err := bookController.PostBookService(&book)
	if err != nil {
		log.Printf("Not able to post Book : %s" , err)
		bookController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	bookController.RespondWithJSON(w, http.StatusCreated, res)
}

func (bookController *BookController) PutBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		bookController.RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	var b Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		bookController.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	res, err := bookController.UpdateBookService(id, &b)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			bookController.RespondWithError(w, http.StatusNotFound, "Book not found")
		default:
			bookController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	bookController.RespondWithJSON(w, http.StatusOK, res)
}

func (bookController *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		bookController.RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	err = bookController.DeleteBookService(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			bookController.RespondWithError(w, http.StatusNotFound, "Book not found")
		default:
			bookController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	bookController.RespondWithJSON(w, http.StatusOK, map[string]string{"result" : "success"})
}