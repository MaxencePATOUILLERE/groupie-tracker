package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		w.Write(data)
	}
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

func locations(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		w.Write(data)
	}
}

func dates(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		w.Write(data)
	}
}

func relation(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		w.Write(data)
	}
}
func main() {
	http.HandleFunc("/api", homePage)
	http.HandleFunc("/api/artists", artists)
	http.HandleFunc("/api/locations", locations)
	http.HandleFunc("/api/dates", dates)
	http.HandleFunc("/api/relation", relation)

	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
