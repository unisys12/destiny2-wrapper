package bungie

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	basePath     string = "https://www.bungie.net/Platform"
	ManifestPath string = "/Destiny2/Manifest/"
)

func Manifest() string {

	// Load Private Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup a client since we need to add a custom header to the request
	client := &http.Client{}

	req, err := http.NewRequest("GET", basePath+ManifestPath, nil)

	if err != nil {
		log.Fatalf("There was an error connecting to the Bugnie API Server: %v", err)
	}

	// Add custom header that includes our API Key
	req.Header.Add("X-API-KEY", os.Getenv("BUNGIE_KEY"))

	// Perform the actual request
	res, err := client.Do(req)

	if err != nil {
		log.Fatalf("There was an error with the Client Request: %v", err)
	}

	// close the request since it succeded
	defer res.Body.Close()

	// Read the request body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("There was an error reading the response: %v", err)
	}

	// Return the request body as a string... for now
	return string(body)

}
