package auth0client

type SigningKey struct {
	Cert    StringOrNil `json:"cert,omitempty"`
	Pkcs7   StringOrNil `json:"pkcs7,omitempty"`
	Subject StringOrNil `json:"subject,omitempty"`
}
