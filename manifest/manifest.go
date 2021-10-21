package manifest

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/unisys12/destiny2-wrapper/bungie"
)

const (
	basePath     string = "https://www.bungie.net/Platform"
	ManifestPath string = "/Destiny2/Manifest/"
)

func init() {
	// Load Private Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ManifestVersion() (string, error) {
	resp, err := http.Get(basePath + ManifestPath)
	if err != nil {
		log.Fatalf("There was an error making a request for the Manifest: %v", err)
	}

	resp.Header.Add("X-API-KEY", os.Getenv("BUNGIE_KEY"))

	defer resp.Body.Close()

	var mResp bungie.ManifestResponse

	if err := json.NewDecoder(resp.Body).Decode(&mResp); err != nil {
		log.Fatalf("There was a problem decoding the JSON Manifest file: %v", err)
	}

	return mResp.Version(), nil
}
