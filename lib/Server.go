package TRC

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Server struct{}

func (serv *Server) Run() {
	FetchArtists()

	http.HandleFunc("/css/", serv.cssHandler)
	http.HandleFunc("/artist/", serv.ArtistHandler)
	http.HandleFunc("/search", serv.SearchHandler)
	http.HandleFunc("/", serv.homeHandler)

	http.ListenAndServe(":8080", nil)
}

func (serv *Server) homeHandler(Writer http.ResponseWriter, Request *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("err ParseFiles index")
	}

	temp.Execute(Writer, Artists)
}

func (serv *Server) cssHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method != http.MethodGet || Request.URL.Path == "/css/" {
		// renderErrorPage(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fileCssServe := http.FileServer(http.Dir("css"))
	http.StripPrefix("/css/", fileCssServe).ServeHTTP(Writer, Request)
}

func (serv *Server) ArtistHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method == "GET" {

		t, err := template.ParseFiles("templates/profile.html")
		if err != nil {
			http.Error(Writer, "500: internal server error", http.StatusInternalServerError)
			return
		}

		ID := string(Request.URL.Path)[len("/artist/"):]
		if Atoi(ID) > len(Artists) {
			// err
			return
		}
		Artists[Atoi(ID)-1].FetchOtherData()

		if err := t.Execute(Writer, Artists[Atoi(ID)-1]); err != nil {
			t, _ = template.ParseFiles("templates/error.html")
			t.Execute(Writer, http.StatusNotFound)
			return
		}

	} else {
		http.Error(Writer, "400: bad request.", http.StatusBadRequest)
	}
}

func (serv *Server) SearchHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method == "POST" {
		var sArtists []Artist
		temp, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(Writer, "500: internal server error", http.StatusInternalServerError)
			return
		}

		Request.ParseForm()

		searchWord := strings.ToLower(Request.FormValue("search"))

		if Location.Index == nil {
			FetchLocations()
		}

		for id := range Artists {
			if artist := Artists[id].Search(searchWord); artist != nil {
				sArtists = append(sArtists, *artist)
			}
		}

		temp.Execute(Writer, sArtists)

	} else {
		http.Error(Writer, "400: bad request.", http.StatusBadRequest)
	}
}
