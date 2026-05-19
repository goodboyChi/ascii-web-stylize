package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"web/asciiArt"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))

type PageData struct {
	Output string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Output: "",
	}
	if r.URL.Path != "/" {
		http.Error(w, "Error 404: You typed a wrong link", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, data)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only a post request is allowed here", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if strings.TrimSpace(text) == ""{
		http.Error(w, "You have to type in a text", http.StatusBadRequest)
		return
	}

	inputText := strings.ReplaceAll(text, "\\n", "\n")
	word := strings.Split(inputText, "\n")
	var ascii strings.Builder

	for _, char := range word {
		if char == "" {
			continue
		}
		ascii.WriteString(asciiArt.Ascii(char, banner))
	}

	data := PageData{
		Output: ascii.String(),
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("template error:", err)
		http.Error(w, "template ERror", http.StatusInternalServerError)
		return
	}
}

func main() {
    http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))


	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii", asciiHandler)
	log.Println("Server listening on: http://localhost:8001")
	http.ListenAndServe(":8001", nil)
}