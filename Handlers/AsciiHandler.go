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
	if len(input) > 1000 {
		http.Error(w, "400 : Bad Request - Input exceeds the maximum allowed length of 1000 characters", http.StatusBadRequest)
		return
	}

	output, status := ascii.PrintAndSplit(input, banner)
	if status != http.StatusOK {
		http.Error(w, output, status)
		return
	}
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
