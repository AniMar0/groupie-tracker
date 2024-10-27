package TRC

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Error struct {
	Err       string
	ErrNumber string
}

func Atoi(num string) int {
	digit := 0
	for _, d := range num {
		if d >= '0' && d <= '9' {
			digit = digit*10 + (int(d) - 48)
		}else{
			return -1
		}
	}
	return digit
}

func renderErrorPage(w http.ResponseWriter, errMsg string, errCode int) {
	var Err Error

	tmpl, tempErr := template.ParseFiles("templates/error.html")
	if tempErr != nil {
		http.Error(w, tempErr.Error(), http.StatusNotFound)
		return
	}
	Err = Error{Err: errMsg, ErrNumber: fmt.Sprintf("%d", errCode)}
	w.WriteHeader(errCode)
	tmpl.Execute(w, Err)
}

func ReplaceAll(searchWord string) string {
	searchWord = strings.ReplaceAll(searchWord, " - artist/band", "")
	searchWord = strings.ReplaceAll(searchWord, " - member", "")
	searchWord = strings.ReplaceAll(searchWord, " - First Album", "")
	searchWord = strings.ReplaceAll(searchWord, " - Creation Date", "")
	searchWord = strings.ReplaceAll(searchWord, " - Location", "")
	return searchWord
}


func (serv *Server) cssHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method != http.MethodGet || Request.URL.Path == "/css/" {
		renderErrorPage(Writer, "Bad Request", http.StatusBadRequest)
		return
	}

	fileCssServe := http.FileServer(http.Dir("css"))
	http.StripPrefix("/css/", fileCssServe).ServeHTTP(Writer, Request)
}

func (serv *Server) jsHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.Method != http.MethodGet || Request.URL.Path == "/js/" {
		renderErrorPage(Writer, "Bad Request", http.StatusBadRequest)
		return
	}

	fileCssServe := http.FileServer(http.Dir("js"))
	http.StripPrefix("/js/", fileCssServe).ServeHTTP(Writer, Request)
}