package main

import (
	"fmt"

	"github.com/unisys12/destiny2-wrapper/bungie"
)

func main() {
	fmt.Println("This is a work in progress")

	manifest, err := bungie.Manifest()
	if err != nil {
		fmt.Printf("%v", err)
	}

	// Returns the version of the current Manifest
	// 97762.21.09.15.0038-0-bnet.40222
	fmt.Println(manifest.Response.Version)

	// Return the MobileGearAssetDataBasePaths Version 0
	// /common/destiny2_content/sqlite/asset/asset_sql_content_daf3f57e18acbfa94afed6379fb91bd9.content
	fmt.Printf("%+v", manifest.Response.MobileGearAssetDataBases[0].Path)
}
