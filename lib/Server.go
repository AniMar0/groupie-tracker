package TRC

import (
	"log"
	"net/http"
	"strconv"
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
		if Atoi(ID) > 51 {
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
	var sArtists []Artist
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(Writer, "500: internal server error", http.StatusInternalServerError)
		return
	}

	Request.ParseForm()

	searchWord := Request.FormValue("search")
	searchWord = strings.ToLower(searchWord)

	FetchLocations()

	for id := range Artists {
		if strings.Contains(strings.ToLower(Artists[id].Name), searchWord) {
			sArtists = append(sArtists, Artists[id])
			continue
		} else if strings.Contains(strings.ToLower(Artists[id].FirstAlbum), searchWord) {
			sArtists = append(sArtists, Artists[id])
			continue
		} else if strings.Contains(strconv.Itoa(Artists[id].CreationDate), searchWord) {
			sArtists = append(sArtists, Artists[id])
			continue
		}

		for membeId := range Artists[id].Members {
			if strings.Contains(strings.ToLower(Artists[id].Members[membeId]), searchWord) {
				sArtists = append(sArtists, Artists[id])
				break
			}
		}

		for locationId := range Location.Index[id].Locations {
			if strings.Contains(strings.ToLower(Location.Index[id].Locations[locationId]), searchWord) {
				sArtists = append(sArtists, Artists[id])
				break
			}
		}
	}
	temp.Execute(Writer, sArtists)
}

// func renderErrorPage(w http.ResponseWriter, errMsg string, errCode int) {
// 	tmpl, tempErr := template.ParseFiles("templates/error.html")
// 	if tempErr != nil {
// 		http.Error(w, tempErr.Error(), http.StatusNotFound)
// 		return
// 	}
// 	Result = Results{Err: errMsg, ErrNumber: fmt.Sprintf("%d", errCode)}
// 	w.WriteHeader(errCode)
// 	tmpl.Execute(w, Result)
// }
