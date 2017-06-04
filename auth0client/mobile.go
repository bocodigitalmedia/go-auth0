package auth0client

type Mobile struct {
	Android *MobileAndroid `json:"android,omitempty"`
	Ios     *MobileIos     `json:"ios,omitempty"`
}

type MobileAndroid struct {
	AppPackageName         interface{} `json:"app_package_name,omitempty"`
	Sha256CertFingerprints *[]string   `json:"sha256_cert_fingerprints,omitempty"`
}

type MobileIos struct {
	TeamId              interface{} `json:"team_id,omitempty"`
	AppBundleIdentifier interface{} `json:"app_bundle_identifier,omitempty"`
}
