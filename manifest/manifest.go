/*
	Package `manifest` consists of methods used to access the Destiny 2 API Manifest.

	Within the manifest, there are the following properties,

	- version

	- mobileAssetContentPath

	- mobileGearAssetDataBases

	- mobileWorldContentPaths

	- jsonWorldContentPaths

	This points to the generated JSON that contains all the Definitions. Each key is
	a locale. The value is a path to the aggregated world definitions (warning: large file!)

	- jsonWorldComponentContentPaths

	This points to the generated JSON that contains all the Definitions. Each key is
	a locale. The value is a dictionary, where the key is a definition type by name,
	and the value is the path to the file for that definition. WARNING: This is unsafe and
	subject to change - do not depend on data in these files staying around long-term.

	- mobileClanBannerDatabasePath

	- mobileGearCDN

	- iconImagePyramidInfo

	Information about the "Image Pyramid" for Destiny icons. Where
	possible, we create smaller versions of Destiny icons. These are found as subfolders under
	the location of the "original/full size" Destiny images, with the same file name and extension
	as the original image itself. (this lets us avoid sending largely redundant path info with
	every entity, at the expense of the smaller versions of the image being less discoverable)
*/
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
	manifestPath string = "/Destiny2/Manifest/"
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
	resp, err := http.Get(basePath + manifestPath)
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
	resp, err := http.Get(basePath + manifestPath)
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
