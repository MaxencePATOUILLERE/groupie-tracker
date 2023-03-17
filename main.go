package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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

func htmlSearchBar(band Artists, location Locations, input string, countrySelect string, nbMembers string, startYear string) string {
	var search string
	var sugestions string
globalLoop:
	for u, v := range band {
		search = ""
		for a, b := range nbMembers {
			buffer, _ := strconv.Atoi(string(b))
			if a == len(nbMembers)-1 && buffer != len(v.Members) {
				continue globalLoop
			} else if buffer == len(v.Members) {
				for _, j := range v.Members {
					search += "<li><a>" + j + "</a></li>"
				}
				break
			} else if buffer != len(v.Members) {
				continue
			}
		}
		countryFound := false
		if countrySelect != "" {
			for concertIndex, concertLocation := range location.Index[u].Locations {
				for position, char := range concertLocation {
					if char == '-' && strings.Compare(concertLocation[position+1:], countrySelect) == 0 {
						search += "<li><a>" + concertLocation + " - " + v.Name + "</a></li>"
						countryFound = true
					} else if !countryFound && concertIndex == len(concertLocation)-1 {
						continue globalLoop
					}
				}
			}
		}

		secondHyphen := false
		for albumDateIndex, albumValue := range v.FirstAlbum {
			if secondHyphen && albumValue == '-' {
				secondHyphen = true
				buffer, _ := strconv.Atoi(v.FirstAlbum[albumDateIndex+1:])
				fmt.Println("buffer : ", buffer)
				secBuffer, _ := strconv.Atoi(startYear)
				if buffer >= secBuffer {
					search += "<li><a>" + v.FirstAlbum + " - " + v.Name + "</a></li>"
				} else {
					continue globalLoop
				}
			}
		}
		buffer, _ := strconv.Atoi(startYear)
		if v.CreationDate >= buffer {
			search += "<li><a>" + strconv.Itoa(v.CreationDate) + " - " + v.Name + "</a></li>"
		} else {
			continue globalLoop
		}

		search += "<li><a>" + v.Name + " - Artist" + "</a></li>"

		sugestions += search
	}
	return sugestions
}

func avoidDouble(input string, country []string) ([]string, bool) {
	for _, coutryWord := range country {
		if len(coutryWord) < len(input) {
			continue
		} else if len(coutryWord) > len(input) {
			continue
		} else {
			for countryletterIndex, countryLetter := range coutryWord {
				if string(input[countryletterIndex]) != string(countryLetter) {
					break
				} else if len(coutryWord)-1 == countryletterIndex {
					return country, true
				}
			}
		}
	}
	country = append(country, input)
	return country, false
}

func htmlDropDown(location Locations) string {
	var htmlString string
	var country []string
	var double bool

	for i := 1; i < 52; i++ {
		for _, v := range location.Index[i].Locations {
			for x, y := range v {
				if y == '-' {
					country, double = avoidDouble(v[x+1:], country)
					if !double {
						htmlString += "<option value=\""
						htmlString += v[x+1:]
						htmlString += "\">"
						htmlString += v[x+1:]
					}
				}
			}
			htmlString += "</option>"
		}
	}
	return htmlString
}

var searchBar, dropDownOptions, input, countrySelect, nbMembers, startYear string

func HomePage(w http.ResponseWriter, r *http.Request) {
	var body string
	artists := getArtist()
	locations := getLocation()

	body += htmlConvert(artists, locations, getDate())

	dropDownOptions = htmlDropDown(locations)
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles("./data/home.html")
		t.Execute(w, map[string]template.HTML{"Body": template.HTML(body), "searchBar": template.HTML(searchBar), "dropDownOptions": template.HTML(dropDownOptions)})
		searchBar = ""
	} else if r.Method == http.MethodPost {
		input = r.FormValue("input")
		countrySelect = r.FormValue("country")
		nbMembers = r.FormValue("nbMembers")
		startYear = r.FormValue("startYear")
		searchBar = "<ul id=\"myUL\">" + htmlSearchBar(artists, locations, input, countrySelect, nbMembers, startYear) + "</ul>"
	}

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
