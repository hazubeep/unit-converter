package main

import (
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

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		lengthTmpl.ExecuteTemplate(w, "layout", nil)
	case http.MethodPost:
		value := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := Result{
			ValueBefore: value,
			ValueAfter:  converter.ConvertLength(v, converter.LengthUnit(from), converter.LengthUnit(to)),
			To:          to,
			From:        from,
		}

		lengthTmpl.ExecuteTemplate(w, "layout", result)
	}
}

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		weightTmpl.ExecuteTemplate(w, "layout", nil)

	case http.MethodPost:
		value := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := Result{
			ValueBefore: value,
			ValueAfter:  converter.ConvertWeight(v, converter.WeightUnit(from), converter.WeightUnit(to)),
			To:          to,
			From:        from,
		}

		weightTmpl.ExecuteTemplate(w, "layout", result)
	}
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		temperatureTmpl.ExecuteTemplate(w, "layout", nil)

	case http.MethodPost:
		value := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		v, err := strconv.ParseFloat(value, 64)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := Result{
			ValueBefore: value,
			ValueAfter:  converter.ConvertTemperature(v, converter.TemperatureUnit(from), converter.TemperatureUnit(to)),
			To:          to,
			From:        from,
		}

		temperatureTmpl.ExecuteTemplate(w, "layout", result)
	}
}
