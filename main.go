package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func artists(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		w.Write(data)
	}
}

func getArtist(id int) Artists {
	var artist Artists
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(id))
	if err != nil {
		// Gérer l'erreur
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&artist)
	return artist
}

func getLocation(id int) Locations {
	var location Locations
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(id))
	if err != nil {
		// Gérer l'erreur
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&location)
	return location
}

func getDate(id int) Dates {
	var date Dates
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(id))
	if err != nil {
		// Gérer l'erreur
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&date)
	return date
}

func main() {
	http.HandleFunc("/", artists)
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
