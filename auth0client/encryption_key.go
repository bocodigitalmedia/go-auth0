package auth0client

type EncryptionKey struct {
	Pub  StringOrNil `json:"pub,omitempty"`
	Cert StringOrNil `json:"cert,omitempty"`
}
