package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

type Generated struct {
	Nick string
}

type WebData struct {
	WebTitle  string
	Nicknames []Generated
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateNickname(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {

	tmpl := template.Must(template.ParseFiles("assets/index.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		data := WebData{
			WebTitle: "Nick Generator",
			Nicknames: []Generated{
				{Nick: generateNickname(10)},
				{Nick: generateNickname(10)},
				{Nick: generateNickname(10)},
				{Nick: generateNickname(10)},
				{Nick: generateNickname(10)},
				{Nick: generateNickname(10)},
				{Nick: generateNickname(10)},
				{Nick: generateNickname(10)},
				{Nick: generateNickname(10)},
				{Nick: generateNickname(10)},
			},
		}

		_ = tmpl.Execute(writer, data)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets/"))))

	_ = http.ListenAndServe(":8080", nil)

}
