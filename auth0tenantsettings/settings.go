package auth0tenantsettings

type Settings struct {
	ChangePassword    *ChangePassword  `json:"change_password,omitempty"`
	GuardianMfaPage   *GuardianMfaPage `json:"guardian_mfa_page,omitempty"`
	DefaultAudience   interface{}      `json:"default_audience,omitempty"`
	DefaultDirectory  interface{}      `json:"default_directory,omitempty"`
	ErrorPage         *ErrorPage       `json:"error_page,omitempty"`
	Flags             *Flags           `json:"flags,omitempty"`
	FriendlyName      interface{}      `json:"friendly_name,omitempty"`
	PictureUrl        interface{}      `json:"picture_url,omitempty"`
	SupportEmail      interface{}      `json:"support_email,omitempty"`
	SupportUrl        interface{}      `json:"support_url,omitempty"`
	AllowedLogoutUrls *[]string        `json:"allowed_logout_urls,omitempty"`
	SessionLifetime   interface{}      `json:"session_lifetime,omitempty"`
}

type ChangePassword struct {
	Enabled interface{} `json:"enabled,omitempty"`
	Html    interface{} `json:"html,omitempty"`
}

type GuardianMfaPage struct {
	Enabled interface{} `json:"enabled,omitempty"`
	Html    interface{} `json:"html, omitempty"`
}

type ErrorPage struct {
	Html        interface{} `json:"html,omitempty"`
	ShowLogLink interface{} `json:"show_log_link,omitempty"`
	Url         interface{} `josn:"url,omitempty"`
}

type Flags struct {
	ChangePwdFlowV1         interface{} `json:"change_pwd_flow_v1,omitempty"`
	EnableApisSection       interface{} `json:"enable_apis_section,omitempty"`
	DisableImpersonation    interface{} `json:"disable_impersonation,omitempty"`
	EnableClientConnections interface{} `json:"enable_client_connections,omitempty"`
	EnablePipeline2         interface{} `json:"enable_pipeline2,omitempty"`
}
