package auth0connection

type GoogleAppsStrategyOptions struct {
	ApiEnableUsers        interface{} `json:"api_enable_users,omitempty"`
	ClientId              interface{} `json:"client_id,omitempty"`
	ClientSecret          interface{} `json:"client_secret,omitempty"`
	DomainAliases         *[]string   `json:"domain_aliases,omitempty"`
	Email                 interface{} `json:"email,omitempty"`
	ExtGroups             interface{} `json:"ext_groups,omitempty"`
	ExtAgreedTerms        interface{} `json:"ext_agreed_terms,omitempty"`
	ExtIsAdmin            interface{} `json:"ext_is_admin,omitempty"`
	ExtIsSuspended        interface{} `json:"ext_is_suspended,omitempty"`
	Global                interface{} `json:"global,omitempty"`
	HandleLoginFromSocial interface{} `json:"handle_login_from_social,omitempty"`
	Profile               interface{} `json:"profile,omitempty"`
	Scope                 *[]string   `json:"scope,omitempty"`
	Status                interface{} `json:"status,omitempty"`
	Strategy              interface{} `json:"strategy,omitempty"`
	TenantDomain          interface{} `json:"tenant_domain,omitempty"`
}
