package auth0client

type JwtConfiguration struct {
	LifetimeInSeconds interface{}             `json:"lifetime_in_seconds,omitempty"`
	Scopes            *map[string]interface{} `json:"scopes,omitempty"`
	Alg               JwtConfigurationAlg     `json:"alg,omitempty"`
}

type JwtConfigurationAlg string

const JwtConfigurationAlgHs256 = JwtConfigurationAlg("HS256")
const JwtConfigurationAlgRs256 = JwtConfigurationAlg("RS256")
