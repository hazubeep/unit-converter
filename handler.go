package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Result struct {
	ValueBefore string
	ValueAfter  float64
	To          string
	From        string
}

func LengthHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/length.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmpl.Execute(w, nil)

	case http.MethodPost:
		value := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		v, _ := strconv.ParseFloat(value, 64)

		convertedResult := ConvertLength(v, Unit(from), Unit(to))

		result := Result{
			ValueBefore: value,
			ValueAfter:  convertedResult,
			To:          to,
			From:        from,
		}

		tmpl.Execute(w, result)
	}
}

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/weight.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmpl.Execute(w, nil)

	case http.MethodPost:
		value := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		v, _ := strconv.ParseFloat(value, 64)

		convertedResult := ConvertWeight(v, WeightUnit(from), WeightUnit(to))

		result := Result{
			ValueBefore: value,
			ValueAfter:  convertedResult,
			To:          to,
			From:        from,
		}

		tmpl.Execute(w, result)
	}
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/temperature.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmpl.Execute(w, nil)

	case http.MethodPost:
		value := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		v, _ := strconv.ParseFloat(value, 64)

		convertedResult := ConvertLength(v, Unit(from), Unit(to))

		result := Result{
			ValueBefore: value,
			ValueAfter:  convertedResult,
			To:          to,
			From:        from,
		}

		tmpl.Execute(w, result)
	}
}
