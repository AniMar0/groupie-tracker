package TRC

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Server struct{}

var Data []string

func (serv *Server) Run() {
	FetchArtists()

	http.HandleFunc("/css/", serv.cssHandler)
	http.HandleFunc("/js/", serv.jsHandler)
	http.HandleFunc("/artist/", serv.ArtistHandler)
	http.HandleFunc("/search", serv.SearchHandler)
	http.HandleFunc("/data", serv.datahandler)
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

func (serv *Server) jsHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method != http.MethodGet || Request.URL.Path == "/css/" {
		// renderErrorPage(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fileCssServe := http.FileServer(http.Dir("js"))
	http.StripPrefix("/js/", fileCssServe).ServeHTTP(Writer, Request)
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

		searchWord := strings.ReplaceAll(Request.FormValue("search"), " - artist/band", "")
		searchWord = strings.ReplaceAll(searchWord, " - member", "")
		searchWord = strings.ReplaceAll(searchWord, " - First Album", "")
		searchWord = strings.ReplaceAll(searchWord, " - Creation Date", "")
		searchWord = strings.ReplaceAll(searchWord, " - Location", "")

		if Location.Index == nil {
			FetchLocations()
		}

		for id := range Artists {
			if artist := Artists[id].Search(strings.ToLower(searchWord)); artist != nil {
				sArtists = append(sArtists, *artist)
			}
		}

		temp.Execute(Writer, sArtists)

	} else {
		http.Error(Writer, "400: bad request.", http.StatusBadRequest)
	}
}

func (serv *Server) datahandler(w http.ResponseWriter, r *http.Request) {
	if Location.Index == nil {
		FetchLocations()
	}

	if Data == nil {
		for id := range Artists {
			Data = append(Data, Artists[id].GetData()...)
		}
	}

	jsonData, err := json.Marshal(Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
