package main

import (
	"fmt"

	"github.com/unisys12/destiny2-wrapper/manifest"
)

func main() {
	fmt.Println("This is a work in progress")

	version, err := manifest.ManifestVersion()
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(version)

	// Return the MobileGearAssetDataBasePaths Version 0
	// /common/destiny2_content/sqlite/asset/asset_sql_content_daf3f57e18acbfa94afed6379fb91bd9.content
	// fmt.Printf("%+v", manifest.Response.MobileGearAssetDataBases[0].Path)
}
