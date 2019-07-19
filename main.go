package main

import (
	"gorilla/mux"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"Projectv2/middlewares/basicauthmiddleware"
	"Projectv2/handlecontrol"
)

func main(){

	y:=mux.NewRouter()
	y.HandleFunc("/", handlecontrol.Home).Methods("GET")
	y.HandleFunc("/topics", handlecontrol.GetTopics).Methods("GET")
	y.Handle("/topics/{name}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlecontrol.GetSubByTopic))).Methods("GET")
	y.Handle("/topics/{name}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlecontrol.CreateTopic))).Methods("POST")
	y.HandleFunc("/topics/{name}/{sub}", handlecontrol.CreateSub).Methods("POST")
	y.Handle("/{name}/{msg}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlecontrol.SendMsg))).Methods("POST")
	y.Handle("/topics/{name}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlecontrol.DeleteTopicByName))).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1010", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(y)))
	}

/*
Documentation:

func NewRouter():
returns a new router instance.

func HandleFunc():
registers a new route with a matcher for the URL path
HandleFunc expects a function

func Handle():
Handle expects a Handler.

func Fatal():
similar to printf()

func ListenAndServe():
ListenAndServe listens on the TCP network address and
then calls Serve with handler to handle requests on incoming connections.

func CORS():
CORS provides Cross-Origin Resource Sharing middleware.

func AllowedHeaders():
adds the provided headers to the list of allowed headers in a CORS request

func AllowedMethods():
AllowedMethods can be used to explicitly allow methods in the Access-Control-Allow-Methods header.

func AllowedOrigins():
AllowedOrigins sets the allowed origins for CORS requests, as used in the 'Allow-Access-Control-Origin'
Passing in a []string{"*"} will allow any domain.

*/
