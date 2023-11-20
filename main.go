package main

// func main() {
// 	reader := getOsReader()
// 	loopFreeBirdTool(reader)
// }

import (
	"flag"
	"log"
)

func main() {

	// Define command-line flags
	city := flag.String("city", "Ho Chi Minh City", "City naming")

	// Parse the command-line flags
	flag.Parse()

	err := loadEnvFile(".env")
	if err != nil {
		log.Fatal(err)
	}

	getWeather(*city) // Handle the flag values
}
