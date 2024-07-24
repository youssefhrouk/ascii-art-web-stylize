package ascii

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Message string
}

// handles requests to the root URL
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 : Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html") // it loads and analyzes the page from the directroy if it doesn't exist we get an error
	if err != nil {
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/" { // handle if the url wasn't valid
		http.NotFound(w, r) // print not found page
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}
}
