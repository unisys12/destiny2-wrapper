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

Import the package using: `"github.com/unisys12/destiny2-wrapper/bungie"` in your import section.

## Usage

To get the current Manifest, which contains paths to all the actually platform data, simply call:

```go
manifest := bungie.Manifest()
```

Currently, this returns something like the following:

```json
{97762.21.09.15.0038-0-bnet.40222 /common/destiny2_content/sqlite/asset/asset_sql_content_daf3f57e18acbfa94afed6379fb91bd9.content /common/destiny2_content/clanbanner/clanbanner_sql_content_16c6ae2603e05f6e11c5894c8b2ec485.content}
```

to get a look at the property keys/labels, just do:

```go
manifest := bungie.Manifest()
fmt.Printf("%+v", manifest.Response)
```

which outputs something like:

```json
{Version:97762.21.09.15.0038-0-bnet.40222 MobileContentPath:/common/destiny2_content/sqlite/asset/asset_sql_content_daf3f57e18acbfa94afed6379fb91bd9.content MobileClanBannerDatabasePath:/common/destiny2_content/clanbanner/clanbanner_sql_content_16c6ae2603e05f6e11c5894c8b2ec485.content}
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
