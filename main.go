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
