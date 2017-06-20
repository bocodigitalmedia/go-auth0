package auth0client

type JwtConfiguration struct {
	LifetimeInSeconds interface{}             `json:"lifetime_in_seconds,omitempty"`
	Scopes            *map[string]interface{} `json:"scopes,omitempty"`
	Alg               interface{}             `json:"alg,omitempty"`
}
