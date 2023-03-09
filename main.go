package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var PORT = 3001

type Artists []struct {
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
type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}
type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

func getArtist() Artists {
	var artist Artists
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		// Gérer l'erreur
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&artist)
	return artist
}

func getLocation() Locations {
	var location Locations
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		// Gérer l'erreur
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&location)
	return location
}

func getDate() Dates {
	var date Dates
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		// Gérer l'erreur
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&date)
	return date
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

func htmlConvert(band Artists, location Locations, date Dates) string {
	var html string
	for u, v := range band {
		html += "<div class=\"card\">\n<div class=\"card__inner\">\n <div class=\"card__face card__face--front\">\n <img src=\""
		html += v.Image
		html += "\" class=\"bandImage\">\n <h2 class=\"BandName\">"
		html += v.Name
		html += "</h2>\n </div>\n <div class=\"card__face card__face--back\">\n <div class=\"card__content\">\n <div class=\"card__header\">\n <img src=\""
		html += v.Image
		html += "\" alt=\"\" class=\"pp\" />\n <h2>"
		html += v.Name
		html += "</h2>\n </div>\n <div class=\"card__body\">"
		html += htmlArtistName(v.Members)
		html += "<p>Creation : " + strconv.Itoa(v.CreationDate) + "</p>"
		html += "<p>First Album : " + v.FirstAlbum + "</p>"
		html += "<p class=\"moreInfo\" onclick=\"popup();\">... More info</p>" + "<div class=\"locationDate\">"
		for i, _ := range location.Index[u].Locations {
			html += "<p class=\"locationDate" + strconv.Itoa(u) + "\">"
			html += htmlConvertLocation(location.Index[u].Locations[i])
			html += htmlConvertDate(date.Index[u].Dates[i])
		}
		html += "\n </div>\n </div>\n </div>\n </div>\n </div>\n </div>"
	}
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
	for u, v := range band {
		search += "<li><a>" + v.Name + "</a></li>"
		for _, j := range v.Members {
			search += "<li><a>" + j + "</a></li>"
		}
		search += "<li><a>" + strconv.Itoa(v.CreationDate) + "</a></li>"
		search += "<li><a>" + v.FirstAlbum + "</a></li>"
		for _, y := range location.Index[u].Locations {
			search += "<li><a>" + y + "</a></li>"
		}
	}
	return search
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	var body, searchBar string
	artists := getArtist()
	locations := getLocation()

	body += htmlConvert(artists, locations, getDate())

	searchBar += "<ul id=\"myUL\">"
	searchBar += htmlSearchBar(artists, locations)
	searchBar += "</ul>"
	t, _ := template.ParseFiles("./data/home.html")
	t.Execute(w, map[string]template.HTML{"Body": template.HTML(body), "searchBar": template.HTML(searchBar)})
}

func main() {
	fileServer := http.FileServer(http.Dir("./data/"))
	http.Handle("/data/", http.StripPrefix("/data/", fileServer))
	http.HandleFunc("/", HomePage)
	fmt.Println("Server on : localhost:3001")
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		return
	}
}
