package main

import (
	"github.com/go-chi/chi"
	"graphql/book"
	"graphql/infrastructure"
	"log"
	"net/http"
	"net/url"
)


func main() {
	routes := chi.NewRouter()
	r := book.RegisterRoutes(routes)
	log.Println("Server ready at 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}
