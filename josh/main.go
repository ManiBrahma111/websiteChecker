package main

import (
	"log"
	"net/http"

	"github.com/mani/handler"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			handler.Getstatus(w)
		}

	case "POST":
		{
			handler.QParam = r.FormValue("name")
			if handler.QParam == "" {
				handler.Postall(w, r)
			} else {
				handler.Postquerry(w, handler.QParam)
			}
		}

	}
}

func main() {

	http.HandleFunc("/websites", requestHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
