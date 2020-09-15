// all Go executables need a main package
package main

// importing packages
// log : loggin status and errors
// net/http : http package for REST API
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User : type of data handled through the API
type User struct {
	userID      string `json:"user___id,omitempty"`
	firstName   string `json:"first___name,omitempty"`
	lastName    string `json:"last___name,omitempty"`
	age         int    `json:"age,omitempty"`
	company     string `json:"comapany,omitempty"`
	designation string `json:"designation,omitempty"`
}

// UserHandelers : handelers for API
type UserHandelers struct {
	store map[string]User
}

// get list of all users
func (h *UserHandelers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// make list of users to store get Data
	users := make([]User, len(h.store))

	i := 0
	for _, user := range h.store {
		users[i] = user
		i++
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	// wrtite response with OK Status
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *UserHandelers) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func (h *UserHandelers) addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func (h *UserHandelers) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func newUserHandelers() *UserHandelers {
	// to be modified later
	file, _ := ioutil.ReadFile("data/uers.json")

	return &UserHandelers{
		store: map[string]User{
			"1": {
				userID:      "U01",
				firstName:   "Kamal",
				lastName:    "Ranga",
				age:         30,
				company:     "zone24x7",
				designation: "Associate Software Engineer",
			},
		},
	}
}

// runnig server on localhost, port : 8080
func main() {
	r := mux.NewRouter()
	userHandelers := newUserHandelers()
	api := r.PathPrefix("/api/v1").Subrouter()

	r.HandleFunc("/get/all", newUserHandelers.getAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/get/{userID}", newUserHandelers.getUser).Methods(http.MethodGet)
	r.HandleFunc("/add", newUserHandelers.addUser).Methods(http.MethodPost)
	r.HandleFunc("/remove/{userID}", newUserHandelers.deleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
