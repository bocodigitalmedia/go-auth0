package auth0client

type SigningKey struct {
	Cert    interface{} `json:"cert,omitempty"`
	Pkcs7   interface{} `json:"pkcs7,omitempty"`
	Subject interface{} `json:"subject,omitempty"`
}
