package TRC

import (
	"encoding/json"
	"net/http"
	"strings"
	"text/template"
)

type Server struct{}

var Data []string

func (serv *Server) Run() {
	FetchArtists()

	http.HandleFunc("/css/", serv.staticFileHandler)
	http.HandleFunc("/js/", serv.staticFileHandler)
	http.HandleFunc("/artist/", serv.ArtistHandler)
	http.HandleFunc("/search", serv.SearchHandler)
	http.HandleFunc("/data", serv.dataHandler)
	http.HandleFunc("/", serv.homeHandler)

	http.ListenAndServe(":8080", nil)
}

func (serv *Server) homeHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method != http.MethodGet || Request.URL.Path != "/" {
		renderErrorPage(Writer, "Not Found", http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderErrorPage(Writer, "Internal Server Error", http.StatusInternalServerError)
	}

	if err := temp.Execute(Writer, Artists); err != nil {
		renderErrorPage(Writer, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (serv *Server) staticFileHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method != http.MethodGet || Request.URL.Path == "/css/" || Request.URL.Path == "/js/" {
		renderErrorPage(Writer, "Bad Request", http.StatusBadRequest)
		return
	}

	if strings.HasPrefix(Request.URL.Path, "/css/") {
		fileCssServe := http.FileServer(http.Dir("css"))
		http.StripPrefix("/css/", fileCssServe).ServeHTTP(Writer, Request)
	} else if strings.HasPrefix(Request.URL.Path, "/js/") {
		fileJsServe := http.FileServer(http.Dir("js"))
		http.StripPrefix("/js/", fileJsServe).ServeHTTP(Writer, Request)
	}
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
	if Atoi(ID) > len(Artists) || Atoi(ID) < 1 {
		renderErrorPage(Writer, "Not Found", http.StatusNotFound)
		return
	}
	Artists[Atoi(ID)-1].FetchOtherData()

	if err := t.Execute(Writer, Artists[Atoi(ID)-1]); err != nil {
		renderErrorPage(Writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (serv *Server) SearchHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method == "POST" {
		var sArtists []Artist
		temp, err := template.ParseFiles("templates/index.html")
		if err != nil {
			renderErrorPage(Writer, "internal server error", http.StatusInternalServerError)
			return
		}

		Request.ParseForm()

		if Location.Index == nil {
			FetchLocations()
		}

		for id := range Artists {
			if artist := Artists[id].Search(strings.ToLower(ReplaceAll(Request.FormValue("search")))); artist != nil {
				sArtists = append(sArtists, *artist)
			}
		}

		if err := temp.Execute(Writer, sArtists); err != nil {
			renderErrorPage(Writer, "Internal Server Error", http.StatusInternalServerError)
		}

	} else {
		renderErrorPage(Writer, "bad request.", http.StatusBadRequest)
	}
}

func (serv *Server) dataHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method == "POST" {
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
			http.Error(Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		Writer.Header().Set("Content-Type", "application/json")
		Writer.Write(jsonData)
	} else {
		renderErrorPage(Writer, "bad request.", http.StatusBadRequest)
	}
}
