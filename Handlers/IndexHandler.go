package ascii

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Message string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 : Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/" { // handle if the url wasn't valid
		http.NotFound(w, r) // print not found page
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
