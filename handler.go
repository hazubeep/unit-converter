package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/hazubeep/unit-converter/converter"
)

type Result struct {
	ValueBefore string
	ValueAfter  float64
	To          string
	From        string
}

type Unit interface {
	converter.LengthUnit | converter.WeightUnit | converter.TemperatureUnit
}

type ConvertFunc[U Unit] func(value float64, from, to U) float64

func MakeHandler[U Unit](tmpl *template.Template, convert ConvertFunc[U]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tmpl.ExecuteTemplate(w, "layout", nil)
		case http.MethodPost:
			value := r.FormValue("value")
			from := r.FormValue("from")
			to := r.FormValue("to")

			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			result := Result{
				ValueBefore: value,
				ValueAfter:  convert(v, U(from), U(to)),
				To:          to,
				From:        from,
			}

			tmpl.ExecuteTemplate(w, "layout", result)
		}
	}
}
