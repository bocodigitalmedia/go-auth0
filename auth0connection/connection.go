package auth0connection

type Connection struct {
	Id string `json:"id,omitempty"`
	Properties
}

type Properties struct {
	Name           interface{} `json:"name,omitempty"`
	Strategy       interface{} `json:"strategy,omitempty"`
	Options        *Options    `json:"options,omitempty"`
	EnabledClients *[]string   `json:"enabled_clients,omitempty"`
	Realms         *[]string   `json:"realms,omitempty"`
}

type Options struct {
	Validation             *map[string]interface{} `json:"validation,omitempty"`
	PasswordPolicy         interface{}             `json:"passwordPolicy,omitempty"`
	PasswordHistory        *map[string]interface{} `json:"password_history,omitempty"`
	PasswordNoPersonalInfo *map[string]interface{} `json:"password_no_personal_info,omitempty"`
	PasswordDictionary     *map[string]interface{} `json:"password_dictionary,omitempty"`
}
