package auth0rule

type Rule struct {
	Id    interface{} `json:"id,omitempty"`
	Stage interface{} `json:"stage,omitempty"`
	Properties
}

type Properties struct {
	Name    interface{} `json:"name,omitempty"`
	Script  interface{} `json:"script,omitempty"`
	Order   interface{} `json:"order,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
}
