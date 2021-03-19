package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)
func main() {
	http.ListenAndServe(":8080", Router().InitRouter())
}
