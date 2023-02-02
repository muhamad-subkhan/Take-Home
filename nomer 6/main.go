package main

import (
	"ShoppingCart/migration"
	"ShoppingCart/router"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment")
		return
	}
	
	migration.Migrate()

	r := mux.NewRouter()

	router.Router(r.PathPrefix("/api").Subrouter())

	fmt.Println("server running localhost:8080")
	http.ListenAndServe("localhost:5000", r)

}