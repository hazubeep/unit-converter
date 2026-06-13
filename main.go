package main

import (
	"html/template"
	"log"
	"net/http"
)

const port = ":8080"

var (
	lengthTmpl      *template.Template
	weightTmpl      *template.Template
	temperatureTmpl *template.Template
)

func main() {
	var err error

	lengthTmpl, err = template.ParseFiles("templates/layout.html", "templates/length.html")
	if err != nil {
		log.Fatal(err)
	}

	weightTmpl, err = template.ParseFiles("templates/layout.html", "templates/weight.html")
	if err != nil {
		log.Fatal(err)
	}

	temperatureTmpl, err = template.ParseFiles("templates/layout.html", "templates/temperature.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", LengthHandler)
	http.HandleFunc("/weight", WeightHandler)
	http.HandleFunc("/temperature", TemperatureHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
