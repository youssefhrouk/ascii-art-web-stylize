package ascii

import (
	"html/template"
	"net/http"

	ascii "ascii/functions"
)

// handles the ASCII art generation requests
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 : Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	input := r.FormValue("input")
	banner := r.FormValue("banner")
	if input == "" || banner == "" {
		http.Error(w, "400 : Bad Request - Input and banner selection are required", http.StatusBadRequest)
		return
	}

	output := ascii.PrintAndSplit(input, banner)

	tmpl, err := template.ParseFiles("templates/ascii-art.html")
	if err != nil {
		http.Error(w, "404 : Not Found - "+err.Error(), http.StatusNotFound)
		return
	}

	data := PageData{
		Message: output,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "500 : Internal Server Error - ", http.StatusInternalServerError)
		return
	}
}
