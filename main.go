package main

import (
	"net/http"

	"github.com/oviniciusfarias/crud-with-go/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
