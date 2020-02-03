package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
)

var (
	apiKey string
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w2, err := owm.NewCurrent("C", "EN", apiKey)
	if err != nil {
		log.Fatal(err)
	}

	w2.CurrentByName("Hong Kong, HK")

	t := fmt.Sprintf("%.f°C\n", w2.Main.Temp)

	w.Write([]byte(t))
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	apiKey = os.Getenv("OWM_API_KEY")

	r := mux.NewRouter()
	r.HandleFunc("/", Handler)

	if port == "" {
		port = ":8000"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	log.Fatal(http.ListenAndServe(port, r))
}
