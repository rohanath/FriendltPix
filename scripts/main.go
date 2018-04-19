package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	Name     string `json:"name" bson:"Name"`
	Price    string `json:"price" bson:"Price"`
}

// OrderController represents the controller for operating on the Order resource
type OrderController struct {
	session *mgo.Session
}

// NewOrderController provides a reference to a OrderController with provided mongo session
func NewOrderController(mgoSession *mgo.Session) *OrderController {
	return &OrderController{mgoSession}
}

func (oc OrderController) GetOrders(w http.ResponseWriter, r *http.Request) {

	var options []Item

	iter := oc.session.DB("test").C("Menu").Find(nil).Iter()
	result := Item{}
	for iter.Next(&result) {
		orders = append(options, result)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&options)
}

func main() {

	r := mux.NewRouter()

	// Get a UserController instance
	oc := NewOrderController(getSession())
	r.Methods("OPTIONS").HandlerFunc(IgnoreOption)

	r.HandleFunc("/v1/starbucks/store3/orders", oc.GetOrders).Methods("GET")
	r.Handle("/", r)
	fmt.Println("serving on port" + GetPort())

	http.ListenAndServe(GetPort(), r)

}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func getSession() (s *mgo.Session) {
	// Connect to local mongodb
	s, _ = mgo.Dial("mongodb:localhost")
	return s
}