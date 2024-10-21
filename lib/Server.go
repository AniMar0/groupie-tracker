package TRC

import (
	"log"
	"net/http"
	"text/template"
)

type Server struct{}

func (serv *Server) Run() {
	FetchArtists()
	http.HandleFunc("/css/", serv.cssHandler)
	http.HandleFunc("/artist/", serv.ArtistHandler)
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

		ID := string(Request.URL.Path)[8:]

		Artists[Atoi(ID)-1].FetchDates()
		Artists[Atoi(ID)-1].FetchLocations()
		Artists[Atoi(ID)-1].FetchRelations()

		if err := t.Execute(Writer, Artists[Atoi(ID)-1]); err != nil {
			t, _ = template.ParseFiles("templates/error.html")
			t.Execute(Writer, http.StatusNotFound)
			return
		}

	} else {
		http.Error(Writer, "400: bad request.", http.StatusBadRequest)
	}
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
