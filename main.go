package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateNickname(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/generate/{nickname}/{nicksCount}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nickname := vars["nickname"]
		nicks := vars["nicksCount"]

		nickLength := len(nickname)

		nicksCount, _ := strconv.Atoi(nicks)

		_, _ = fmt.Fprintln(w, "")

		for i := 0; i < nicksCount; i++ {
			_, _ = fmt.Fprintln(w, nickname+"_"+generateNickname(15-nickLength))
			_, _ = fmt.Fprintln(w, "")
		}
	})

	_ = http.ListenAndServe(":8080", r)

}
