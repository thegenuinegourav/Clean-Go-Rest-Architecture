package main

import (
	"github.com/Crud/Book"
	common "github.com/Crud/Common"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

type kernel struct {}

type IKernel interface {
	InjectDB()
	InjectBookController() Book.BookController
}

var(
	k				*kernel
	kernelOnce 	sync.Once
	db 				*gorm.DB
	err 			error
)

// Singleton
func Kernel() IKernel {
	if k == nil {
		kernelOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

func (k *kernel) InjectDB() {
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Football?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Database connection failed : ", err)
	}else {
		log.Println("Database connection established!")
	}
	db.AutoMigrate(&Book.Book{})
}

func (k *kernel) InjectBookController() Book.BookController {
	bookRepository := &Book.BookRepository{db}
	bookService := &Book.BookService{bookRepository}
	bookController := &Book.BookController{bookService, &common.CommonController{}}
	return *bookController
}
