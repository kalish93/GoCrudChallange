package main

import (
	"encoding/json"
	"log"
	"net/http"
    "github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Person struct {
	Id    string   `json:"id"`;
	Name  string   `json:"name"`;
	Age   int      `json:"age"`;
	Hobbies  []string  `json:"hobbies"`;
}

var persons []Person

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/person", createPersonHandler).Methods("POST")
    r.HandleFunc("/person", getPersonsHandler).Methods("GET")
    r.HandleFunc("/person/{id}", getPersonHandler).Methods("GET")
    r.HandleFunc("/person/{id}", updatePersonHandler).Methods("PUT")
    r.HandleFunc("/person/{id}", deletePersonHandler).Methods("DELETE")

    log.Println("Server listening on :3000")
    log.Fatal(http.ListenAndServe(":3000", r))
}

func getPersonsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(persons);
}

func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.Id = uuid.New().String()
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)
}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range persons {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func updatePersonHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for index, item := range persons {
		if item.Id == params["id"] {
			var updatedPerson Person
			_ = json.NewDecoder(r.Body).Decode(&updatedPerson)
			updatedPerson.Id = persons[index].Id
			persons[index] = updatedPerson
			json.NewEncoder(w).Encode(updatedPerson)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{});
}

func deletePersonHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range persons {
		if item.Id == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)
}