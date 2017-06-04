package auth0client

type EncryptionKey struct {
	Pub  interface{} `json:"pub,omitempty"`
	Cert interface{} `json:"cert,omitempty"`
}
