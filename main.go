package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/hazubeep/unit-converter/converter"
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

	http.HandleFunc("/", MakeHandler(lengthTmpl, converter.ConvertLength))
	http.HandleFunc("/weight", MakeHandler(weightTmpl, converter.ConvertWeight))
	http.HandleFunc("/temperature", MakeHandler(temperatureTmpl, converter.ConvertTemperature))

	fmt.Printf("server run at port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
