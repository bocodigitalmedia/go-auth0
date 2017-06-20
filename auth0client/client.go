package auth0client

type Client struct {
	ClientId     interface{}   `json:"client_id,omitempty"`
	ClientSecret interface{}   `json:"client_secret",omitempty"`
	IsFirstParty interface{}   `json:"is_first_party,omitempty"`
	SigningKeys  *[]SigningKey `json:"signing_keys,omitempty"`
	Global       interface{}   `json:"global,omitempty"`
	Tenant       interface{}   `json:"tenant,omitempty"`

	Properties
}

type Properties struct {
	Name                    interface{}             `json:"name,omitempty"`
	Description             interface{}             `json:"description,omitempty"`
	LogoUri                 interface{}             `json:"logo_uri,omitempty"`
	Callbacks               *[]interface{}          `json:"callbacks,omitempty"`
	AllowedOrigins          *[]interface{}          `json:"allowed_origins,omitempty"`
	ClientAliases           *[]interface{}          `json:"client_aliases,omitempty"`
	AllowedClients          *[]interface{}          `json:"allowed_clients,omitempty"`
	AllowedLogoutUrls       *[]interface{}          `json:"allowed_logout_urls,omitempty"`
	GrantTypes              *[]interface{}          `json:"grant_types,omitempty"`
	TokenEndpointAuthMethod interface{}             `json:"token_endpoint_auth_method,omitempty"`
	AppType                 interface{}             `json:"app_type,omitempty"`
	OidcConformant          interface{}             `json:"oidc_conformant,omitempty"`
	JwtConfiguration        *JwtConfiguration       `json:"jwt_configuration,omitempty"`
	EncryptionKey           *EncryptionKey          `json:"encryption_key,omitempty"`
	Sso                     interface{}             `json:"sso,omitempty"`
	CrossOriginAuth         interface{}             `json:"cross_origin_auth,omitempty"`
	CrossOriginLoc          interface{}             `json:"cross_origin_loc,omitempty"`
	SsoDisabled             interface{}             `json:"sso_disabled,omitempty"`
	CustomLoginPageOn       interface{}             `json:"custom_login_page_on,omitempty"`
	CustomLoginPage         interface{}             `json:"custom_login_page,omitempty"`
	CustomLoginPagePreview  interface{}             `json:"custom_login_page_preview,omitempty"`
	FormTemplate            interface{}             `json:"form_template,omitempty"`
	IsHerokuApp             interface{}             `json:"is_heroku_app,omitempty"`
	Addons                  *Addons                 `json:"addons,omitempty"`
	ClientMetadata          *map[string]interface{} `json:"client_metadata,omitempty"`
	Mobile                  *Mobile                 `json:"mobile,omitempty"`
}
