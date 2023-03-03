package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var PORT = 8080

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

func htmlArtistName(members []string) string {
	var htmlmembers string
	htmlmembers += "<p>Members : "
	for _, v := range members {
		if v == members[len(members)-1] {
			break
		}
		htmlmembers += v + ", "
	}
	htmlmembers += members[len(members)-1]
	htmlmembers += "</p>"
	return htmlmembers
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
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

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
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

func htmlConvert(band Artists, location Locations, date Dates) string {
	var html string
	html += "<div class=\"card\">\n<div class=\"card__inner\">\n <div class=\"card__face card__face--front\">\n <img src=\""
	html += band.Image
	html += "\" class=\"bandImage\">\n <h2 class=\"BandName\">"
	html += band.Name
	html += "</h2>\n </div>\n <div class=\"card__face card__face--back\">\n <div class=\"card__content\">\n <div class=\"card__header\">\n <img src=\""
	html += band.Image
	html += "\" alt=\"\" class=\"pp\" />\n <h2>"
	html += band.Name
	html += "</h2>\n </div>\n <div class=\"card__body\">"
	html += htmlArtistName(band.Members)
	html += "<p>Creation : " + strconv.Itoa(band.CreationDate) + "</p>"
	html += "<p>First Album : " + band.FirstAlbum + "</p>"
	html += "<p class=\"moreInfo\" onclick=\"popup();\">... More info</p>" + "<div class=\"locationDate\">"
	for i := 0; i < len(location.Locations); i++ {
		html += "<p class=\"locationDate" + strconv.Itoa(i) + "\">"
		html += htmlConvertLocation(location.Locations[i])
		html += htmlConvertDate(date.Dates[i])
	}
	html += "\n </div>\n </div>\n </div>\n </div>\n </div>\n </div>"
	return html
}

func htmlConvertLocation(location string) string {
	var htmlLocation string
	htmlLocation += "Location: " + location + ", "
	return htmlLocation
}

func htmlConvertDate(date string) string {
	var htmlDate string
	htmlDate += "Date: " + date + "</p>"
	return htmlDate
}

func htmlSearchBar(band Artists, location Locations) string {
	var search string
	search += "<li><a>" + band.Name + "</a></li>"
	for _, v := range band.Members {
		search += "<li><a>" + v + "</a></li>"
	}
	search += "<li><a>" + strconv.Itoa(band.CreationDate) + "</a></li>"
	search += "<li><a>" + band.FirstAlbum + "</a></li>"
	for _, v := range location.Locations {
		search += "<li><a>" + v + "</a></li>"
	}
	return search
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	var body, searchBar string
	searchBar += "<ul id=\"myUL\">"
	for i := 1; i <= 52; i++ {
		artistButter := getArtist(i)
		locationButter := getLocation(i)
		dateButter := getDate(i)
		body += htmlConvert(artistButter, locationButter, dateButter)
		searchBar += htmlSearchBar(artistButter, locationButter)
	}
	searchBar += "</ul>"
	t, _ := template.ParseFiles("./data/home.html")
	t.Execute(w, map[string]template.HTML{"Body": template.HTML(body), "searchBar": template.HTML(searchBar)})
}

func main() {
	fileServer := http.FileServer(http.Dir("./data/"))
	http.Handle("/data/", http.StripPrefix("/data/", fileServer))
	http.HandleFunc("/", HomePage)
	fmt.Println("Server on : localhost:8080")
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		return
	}
}
