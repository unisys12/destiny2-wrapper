package main

import (
	"fmt"

	"github.com/unisys12/destiny2-wrapper/bungie"
)

func main() {
	fmt.Println("This is a work in progress")

	manifest := bungie.Manifest()

	fmt.Println(manifest.Response.Version)
}
