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

// ManifestVersion
//
// Returns a string of the current version of the Manifest from
// the Destiny 2 API. Will and can be used internally and externally
// to check for Manifest updates, which happen on Tuesdays.
// When the Manifest updates, the game updates. This doesn't mean
// new content has been added. It only means that a new build of the
// game was pushed out.
//
// 		version, err := manifest.ManifestVersion()
//		if err != nil {
//			fmt.Printf("%v", err)
//		}
//		fmt.Println(version)
//
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

// MobileAsssetContentPath returns a string representing the latest
// path to the MobileAsset .content file, which is a SQLite database
// that contains... Ya know, I have no idea. It's currently not
// documented by Bungie, so knock yourself out.
//
//		macp, err := manifest.MobileAsssetContentPath()
//		if err != nil {
//			fmt.Printf("%v", err)
//		}
//		fmt.Println(macp)
//
func MobileAsssetContentPath() (string, error) {
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

	return mResp.MobileAssetContentPath(), nil
}
