package auth0client

type Mobile struct {
	Android *MobileAndroid `json:"android,omitempty"`
	Ios     *MobileIos     `json:"ios,omitempty"`
}

type MobileAndroid struct {
	AppPackageName         StringOrNil `json:"app_package_name,omitempty"`
	Sha256CertFingerprints *[]string   `json:"sha256_cert_fingerprints,omitempty"`
}

type MobileIos struct {
	TeamId              StringOrNil `json:"team_id,omitempty"`
	AppBundleIdentifier StringOrNil `json:"app_bundle_identifier,omitempty"`
}
