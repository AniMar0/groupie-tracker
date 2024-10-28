package TRC

import (
	"net/http"
	"strings"
	"text/template"
)

type Server struct{}

var Data []string

func (serv *Server) Run() {
	FetchArtists()
	FetchLocations()
	Append()

	mm()

	http.HandleFunc("/css/", serv.staticFileHandler)
	http.HandleFunc("/artist/", serv.ArtistHandler)
	http.HandleFunc("/search", serv.SearchHandler)
	http.HandleFunc("/", serv.homeHandler)

	http.ListenAndServe(":8082", nil)
}

func (serv *Server) homeHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method != http.MethodGet || Request.URL.Path != "/" {
		renderErrorPage(Writer, "Not Found", http.StatusNotFound)
		return
	}
	// Alle.SArtists = Alle.Artists
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderErrorPage(Writer, "Internal Server Error", http.StatusInternalServerError)
	}

	if err := temp.Execute(Writer, Alle); err != nil {
		renderErrorPage(Writer, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (serv *Server) staticFileHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method != http.MethodGet || Request.URL.Path == "/css/" {
		renderErrorPage(Writer, "Bad Request", http.StatusBadRequest)
		return
	}

	fileCssServe := http.FileServer(http.Dir("css"))
	http.StripPrefix("/css/", fileCssServe).ServeHTTP(Writer, Request)
}

func (serv *Server) ArtistHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.URL.Path == "/artist/" || Request.Method != "GET" {
		renderErrorPage(Writer, "bad request.", http.StatusBadRequest)
		return
	}

	t, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		renderErrorPage(Writer, "internal server error", http.StatusInternalServerError)
		return
	}

	ID := string(Request.URL.Path)[len("/artist/"):]
	if Atoi(ID) > len(Alle.Artists) || Atoi(ID) < 1 {
		renderErrorPage(Writer, "Not Found", http.StatusNotFound)
		return
	}
	Alle.Artists[Atoi(ID)-1].FetchOtherData()

	if err := t.Execute(Writer, Alle.Artists[Atoi(ID)-1]); err != nil {
		renderErrorPage(Writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (serv *Server) SearchHandler(Writer http.ResponseWriter, Request *http.Request) {
	Alle.SArtists = nil
	if Request.Method == "POST" {
		temp, err := template.ParseFiles("templates/search.html")
		if err != nil {
			renderErrorPage(Writer, "internal server error", http.StatusInternalServerError)
			return
		}

		Request.ParseForm()

		for id := range Alle.Artists {
			if artist := Alle.Artists[id].Search(strings.ToLower(ReplaceAll(Request.FormValue("search")))); artist != nil {
				Alle.SArtists = append(Alle.SArtists, *artist)
			}
		}

		if err := temp.Execute(Writer, Alle); err != nil {
			renderErrorPage(Writer, "Internal Server Error", http.StatusInternalServerError)
		}

	} else {
		renderErrorPage(Writer, "bad request.", http.StatusBadRequest)
	}
}
