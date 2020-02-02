package main

import (
	"fmt"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("OWM_API_KEY")

	w, err := owm.NewCurrent("C", "EN", apiKey)
	if err != nil {
		log.Fatal(err)
	}

	w.CurrentByName("Hong Kong, HK")

	fmt.Printf("%.f°C\n", w.Main.Temp)
}
