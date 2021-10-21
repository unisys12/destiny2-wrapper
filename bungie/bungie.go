package bungie

type ManifestResponse struct {
	Response        ResponseProp
	ErrorCode       int32
	ThrottleSeconds int32
	ErrorStatus     string
	Message         string
	// MessageData        string // need to work on this
	DetailedErrorTrace string
}

type ResponseProp struct {
	Version                  string `json:"version"`
	MobileAssetContentPath   string `json:"mobileAssetContentPath"`
	MobileGearAssetDataBases MobileGearAssetDataBasesResponse
	// mobileWorldContentPaths        interface{}
	// jsonWorldContentPaths          interface{}
	// jsonWorldComponentContentPaths interface{}
	MobileClanBannerDatabasePath string `json:"mobileClanBannerDatabasePath"`
	// mobileGearCDN                  interface{}
	// iconImagePyramidInfo           interface{}
}

type MobileGearAssetDataBasesResponse []struct {
	Version int    `json:"version"`
	Path    string `json:"path"`
}

func (m ManifestResponse) Version() string {
	return m.Response.Version
}
