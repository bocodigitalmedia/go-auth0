package auth0connection

import "encoding/json"

type Connection struct {
	Id interface{} `json:"id,omitempty"`
	Properties
}

type Properties struct {
	Name           interface{}    `json:"name,omitempty"`
	Strategy       interface{}    `json:"strategy,omitempty"`
	Options        interface{}    `json:"options,omitempty"`
	EnabledClients *[]interface{} `json:"enabled_clients,omitempty"`
	Realms         *[]interface{} `json:"realms,omitempty"`
}

func (p *Properties) GetOptions() (interface{}, error) {
	if bytes, err := json.Marshal(p.Options); err != nil {
		return nil, err
	} else {
		switch p.Strategy {
		case "google-apps":
			result := new(GoogleAppsStrategyOptions)
			err := json.Unmarshal(bytes, result)
			return result, err
		}
	}
	return nil, nil
}

//
// type Options struct {
// 	Validation             *map[string]interface{} `json:"validation,omitempty"`
// 	PasswordPolicy         interface{}             `json:"passwordPolicy,omitempty"`
// 	PasswordHistory        *map[string]interface{} `json:"password_history,omitempty"`
// 	PasswordNoPersonalInfo *map[string]interface{} `json:"password_no_personal_info,omitempty"`
// 	PasswordDictionary     *map[string]interface{} `json:"password_dictionary,omitempty"`
// }
