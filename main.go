package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", LengthHandler)
	http.HandleFunc("/weight", WeightHandler)
	http.HandleFunc("/temperature", TemperatureHandler)

	http.ListenAndServe(":8080", nil)
}
