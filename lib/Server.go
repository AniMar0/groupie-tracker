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

func (serv *Server) ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		t, err := template.ParseFiles("templates/profile.html")
		if err != nil {
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}

		ID := string(r.URL.Path)[8:]

		for i, v := range Artists {
			if v.ID == Atoi(ID) {
				Artists[i].Relations = Artists[i].OtherDatesLocationsInfos.DatesLocations
				t.Execute(w, Artists[i])
				return
			}
		}

		t, _ = template.ParseFiles("templates/error.html")
		t.Execute(w, http.StatusNotFound)
		return
	} else {
		http.Error(w, "400: bad request.", http.StatusBadRequest)
	}
}
