package auth0client

type Addons struct {
	Aws                  *Addon `json:"aws,omitempty"`
	Firebase             *Addon `json:"firebase,omitempty"`
	AzureBlob            *Addon `json:"azure_blob,omitempty"`
	AzureSb              *Addon `json:"azure_sb,omitempty"`
	Rms                  *Addon `json:"rms,omitempty"`
	MsCrm                *Addon `json:"mscrm,omitempty"`
	Slack                *Addon `json:"slack,omitempty"`
	Box                  *Addon `json:"box,omitempty"`
	CloudBees            *Addon `json:"cloudbees,omitempty"`
	Concur               *Addon `json:"concur,omitempty"`
	Dropbox              *Addon `json:"dropbox,omitempty"`
	EchoSign             *Addon `json:"echosign,omitempty"`
	Egnyte               *Addon `json:"egnyte,omitempty"`
	NewRelic             *Addon `json:"newrelic,omitempty"`
	Office365            *Addon `json:"office365,omitempty"`
	Salesforce           *Addon `json:"salesforce,omitempty"`
	SalesforceApi        *Addon `json:"salesforce_api,omitempty"`
	SalesforceSandboxApi *Addon `json:"salesforce_sandbox_api,omitempty"`
	SamlP                *Addon `json:"samlp,omitempty"`
	Layer                *Addon `json:"layer,omitempty"`
	SapApi               *Addon `json:"sap_api,omitempty"`
	SharePoint           *Addon `json:"sharepoint,omitempty"`
	SpringCm             *Addon `json:"springcm,omitempty"`
	Wams                 *Addon `json:"wams,omitempty"`
	WsFed                *Addon `json:"wsfed,omitempty"`
	Zendesk              *Addon `json:"zendesk,omitempty"`
	Zoom                 *Addon `json:"zoom,omitempty"`
}

type Addon map[string]interface{}

func NewAddonAws() Addon {
	return Addon{}
}

func NewAddonFirebase() Addon {
	return Addon{
		"secret":              nil,
		"client_email":        nil,
		"private_key":         nil,
		"private_key_id":      nil,
		"lifetime_in_seconds": nil,
	}
}
