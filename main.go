package main

import (
	"html/template"
	"net/http"
)

type WebData struct {
	WebTitle string
}

func main() {


	tmpl := template.Must(template.ParseFiles("assets/index.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		data := WebData{
			WebTitle: "Nick Generator",
		}

		_ = tmpl.Execute(writer, data)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets/"))))

	_ = http.ListenAndServe(":8080", nil)

}
