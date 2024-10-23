package TRC

func Atoi(num string) int {
	digit := 0
	for _, d := range num {
		if d >= '0' && d <= '9' {
			digit = digit*10 + (int(d) - 48)
		}
	}
	return digit
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
