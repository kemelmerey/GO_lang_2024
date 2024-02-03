package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type K_dramas struct{
	Title        string `json:"Title"`
	Genre       string `json:"Genre"`
}
type kdramalist []K_dramas

type Response struct {
	Persons []Person `json:"friends"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type SingleFriend []Friend

type Friend struct {
	Id        string    `json:"Id"`
	FirstName string 	`json:"First_name"`
	LastName  string 	`json:"Last_name"`
	Age 	  string 	`json:"Age"`
	Hobby  	  string 	`json:"Hobby"`
}

func main(){
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health_check", HealthCheck).Methods("GET")
	router.HandleFunc("/friends", Friends).Methods("GET")
	router.HandleFunc("/friends/{FirstName}", getOneFriend).Methods("GET")
	router.HandleFunc("/", homeLink)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}

var dramas = kdramalist{
	{
		Title: "The Mouse",
		Genre: "Crime thriller",
	},
	{
		Title: "The happiness",
		Genre: "apocalyptic thriller",
	},
	{
		Title: "Flower of Evil",
		Genre: "Romance",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to my page!. I like watching k-dramas. Here is the list.")
	json.NewEncoder(w).Encode(dramas)
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Description: This app is used to give you information about me and my friends. Welcome to my app;)\nAuthor: Kemel Merey")
}

func getOneFriend(w http.ResponseWriter, r *http.Request){
	frName := mux.Vars(r)["FirstName"]

	for _, singleFriend := range fr {
		if singleFriend.FirstName == frName {
			json.NewEncoder(w).Encode(singleFriend)
		}
	}
}

func Friends(w http.ResponseWriter, r *http.Request){
	log.Println("entering friends end point")
	var response Response 
	friends := prepareResponse()

	response.Persons = friends

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return 
	}

	w.Write(jsonResponse)
}

var fr = SingleFriend{
	{
		Id: "1",       
		FirstName: "Merey", 
		LastName: "Kemelbay",
		Age: "19",
		Hobby: "Play games",
	},
	{
		Id: "2",       
		FirstName: "Dariga", 
		LastName: "Orazbay",
		Age: "19",
		Hobby: "Drawing",
	},
	{
		Id: "3",       
		FirstName: "Ulpash", 
		LastName: "Kuanyshbek",
		Age: "19",
		Hobby: "reading books",
	},
}

func prepareResponse() []Person{
	var friendss []Person

	var person Person
	person.Id = 1
	person.FirstName = "Merey"
	person.LastName = "Kemelbay"
	friendss = append(friendss, person)

	person.Id = 2
	person.FirstName = "Dariga"
	person.LastName = "Orazbay"
	friendss = append(friendss, person)

	person.Id = 3
	person.FirstName = "Ulpash"
	person.LastName = "Kuanyshbek"
	friendss = append(friendss, person)
	return friendss
}