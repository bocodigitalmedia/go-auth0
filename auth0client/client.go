package auth0client

type Client struct {
	ClientId     StringOrNil   `json:"client_id,omitempty"`
	ClientSecret StringOrNil   `json:"client_secret",omitempty"`
	IsFirstParty BoolOrNil     `json:"is_first_party,omitempty"`
	SigningKeys  *[]SigningKey `json:"signing_keys,omitempty"`
	Global       BoolOrNil     `json:"global,omitempty"`
	Tenant       StringOrNil   `json:"tenant,omitempty"`

	*Properties
}

type Properties struct {
	Name                    StringOrNil             `json:"name,omitempty"`
	Description             StringOrNil             `json:"description,omitempty"`
	LogoUri                 StringOrNil             `json:"logo_uri,omitempty"`
	Callbacks               *[]string               `json:"callbacks,omitempty"`
	AllowedOrigins          *[]string               `json:"allowed_origins,omitempty"`
	ClientAliases           *[]string               `json:"client_aliases,omitempty"`
	AllowedClients          *[]string               `json:"allowed_clients,omitempty"`
	AllowedLogoutUrls       *[]string               `json:"allowed_logout_urls,omitempty"`
	GrantTypes              *[]string               `json:"grant_types",omitempty`
	TokenEndpointAuthMethod TokenEndpointAuthMethod `json:"token_endpoint_auth_method,omitempty"`
	AppType                 AppType                 `json:"app_type,omitempty"`
	OidcConformant          BoolOrNil               `json:"oidc_conformant,omitempty"`
	JwtConfiguration        *JwtConfiguration       `json:"jwt_configuration,omitempty"`
	EncryptionKey           *EncryptionKey          `json:"encryption_key,omitempty"`
	Sso                     BoolOrNil               `json:"sso,omitempty"`
	CrossOriginAuth         BoolOrNil               `json:"cross_origin_auth,omitempty"`
	CrossOriginLoc          StringOrNil             `json:"cross_origin_loc,omitempty"`
	SsoDisabled             BoolOrNil               `json:"sso_disabled,omitempty"`
	CustomLoginPageOn       BoolOrNil               `json:"custom_login_page_on,omitempty"`
	CustomLoginPage         StringOrNil             `json:"custom_login_page,omitempty"`
	CustomLoginPagePreview  StringOrNil             `json:"custom_login_page_preview,omitempty"`
	FormTemplate            StringOrNil             `json:"form_template,omitempty"`
	IsHerokuApp             BoolOrNil               `json:"is_heroku_app,omitempty"`
	Addons                  *Addons                 `json:"addons,omitempty"`
	ClientMetadata          *ClientMetadata         `json:"client_metadata,omitempty"`
	Mobile                  *Mobile                 `json:"mobile,omitempty"`
}

type ClientMetadata map[string]interface{}
