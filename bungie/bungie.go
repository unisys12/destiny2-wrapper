package bungie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	basePath     string = "https://www.bungie.net/Platform"
	ManifestPath string = "/Destiny2/Manifest/"
)

type ManifestResponse struct {
	Response        ResponseProp
	ErrorCode       int32
	ThrottleSeconds int32
	ErrorStatus     string
	Message         string
	// MessageData        string	// need to work on this
	DetailedErrorTrace string
}

type ResponseProp struct {
	Version           string `json:"version"`
	MobileContentPath string `json:"mobileAssetContentPath"`
	// mobileGearAssetDataBases       interface{}
	// mobileWorldContentPaths        interface{}
	// jsonWorldContentPaths          interface{}
	// jsonWorldComponentContentPaths interface{}
	MobileClanBannerDatabasePath string `json:"mobileClanBannerDatabasePath"`
	// mobileGearCDN                  interface{}
	// iconImagePyramidInfo           interface{}
}

type MobileGearAssetDataBasesResponse struct {
}

func init() {
	// Load Private Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Manifest() *ManifestResponse {
	// Setup a client since we need to add a custom header to the request
	client := &http.Client{}

	req, err := http.NewRequest("GET", basePath+ManifestPath, nil)

	if err != nil {
		log.Fatalf("There was an error connecting to the Bugnie API Server: %v", err)
	}

	// // Add custom header that includes our API Key
	req.Header.Add("X-API-KEY", os.Getenv("BUNGIE_KEY"))

	// // Perform the actual request
	res, clienterr := client.Do(req)

	if clienterr != nil {
		log.Fatalf("There was an error with the Client Request: %v", clienterr)
	}

	// // close the request since it succeded
	defer res.Body.Close()

	// // Read the request body
	body, readerr := ioutil.ReadAll(res.Body)
	if readerr != nil {
		log.Fatalf("There was an error reading the response: %v", readerr)
	}

	var manifestResponse ManifestResponse

	jsonErr := json.Unmarshal(body, &manifestResponse)
	if jsonErr != nil {
		log.Fatalf("There was an issue while decoding the json response: %v", jsonErr)
	}
	// fmt.Printf("%+v", manifestResponse)
	return &manifestResponse
}
