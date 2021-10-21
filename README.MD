# Destiny2 Wrapper (WIP)

Basic Go Module that exposes methods that are used to communicate with the Bungie Destiny 2 API. At least, that's the plan anyway.

## Setup

### Environment

You will need a `.env` file to hold your required `X-API-KEY`. This key is your developer key that you acquire from the [Bungie Dev Portal](https://www.bungie.net/en/Application). Simply log into Bungie.net account and create a new App. You will get a key and you should paste that key into your `.env` file like so:

```
BUNGIE_KEY=<your super secret key goes here>
```

It is required for the package to work, so...

### Import

Import the package using: `"github.com/unisys12/destiny2-wrapper/manifest"` in your import section.

## Usage

To get the current Manifest, which contains paths to all the actually platform data, simply call:

```go
version, err := manifest.ManifestVersion()
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(version)
```

Currently, this returns something like the following:

```shell
97762.21.09.15.0038-0-bnet.40222
```


## Wrapper Responses

All responses are laid out in a Struct that can be found in `bungie.go`:

```go
// All HTTP requests are returned in this format
type ManifestResponse struct {
	Response        ResponseProp
	ErrorCode       int32
	ThrottleSeconds int32
	ErrorStatus     string
	Message         string
	// MessageData        string	// need to work on this
	DetailedErrorTrace string
}
```

What we are most concerned with, when everything works, is the `Response` property, which I have also laid out in a Struct:

```go
// Response Struct
type ResponseProp struct {
	Version           		string `json:"version"`
	MobileContentPath 		string `json:"mobileAssetContentPath"`
	MobileGearAssetDataBases MobileGearAssetDataBasesResponse
	// mobileWorldContentPaths        interface{}
	// jsonWorldContentPaths          interface{}
	// jsonWorldComponentContentPaths interface{}
	MobileClanBannerDatabasePath string `json:"mobileClanBannerDatabasePath"`
	// mobileGearCDN                  interface{}
	// iconImagePyramidInfo           interface{}
}
```

I'm currently working on getting the properties that are commented out supported. As they are, I will uncomment them of course.
