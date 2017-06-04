package auth0resourceserver

type ResourceServer struct {
	Properties
}

type Properties struct {
	Name                 interface{} `json:"name,omitempty"`
	Identifier           interface{} `json:"identifier,omitempty"`
	Scopes               *[]Scope    `json:"scopes,omitempty"`
	SigningAlg           interface{} `json:"signing_alg,omitempty"`
	SigningSecret        interface{} `json:"signing_secret,omitempty"`
	AllowOfflineAccess   interface{} `json:"allow_offline_access,omitempty"`
	TokenLifetime        interface{} `json:"token_lifetime,omitempty"`
	VerificationKey      interface{} `json:"verificationKey,omitempty"`
	VerificationLocation interface{} `json:"verificationLocation,omitempty"`

	SkipConsentForVerifiableFirstPartyClients interface{} `json:"skip_consent_for_verifiable_first_party_clients,omitempty"`
}

type Scope struct {
	Description interface{} `json:"description,omitempty"`
	Value       interface{} `json:"value,omitempty"`
}
