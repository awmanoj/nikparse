package nikparse

import (
	"embed"
	"encoding/json"
	"log"
)

//go:embed assets/data.txt
var dataFile embed.FS

type Data map[string]map[string]string
var GeoData Data

func init() {
	// Read the embedded file
	byteValue, err := dataFile.ReadFile("assets/data.txt")
	if err != nil {
		log.Fatalf("err failed to read embedded file: %s", err)
	}

	// Parse the JSON data
	err = json.Unmarshal(byteValue, &GeoData)
	if err != nil {
		log.Fatalf("err failed to parse JSON: %s", err)
	}
}
