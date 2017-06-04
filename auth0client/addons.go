package auth0client

type Addons struct {
	Aws      *AddonAws      `json:"aws,omitempty"`
	Firebase *AddonFirebase `json:"firebase,omitempty"`

	AzureBlob            *map[string]interface{} `json:"azure_blob,omitempty"`
	AzureSb              *map[string]interface{} `json:"azure_sb,omitempty"`
	Rms                  *map[string]interface{} `json:"rms,omitempty"`
	MsCrm                *map[string]interface{} `json:"mscrm,omitempty"`
	Slack                *map[string]interface{} `json:"slack,omitempty"`
	Box                  *map[string]interface{} `json:"box,omitempty"`
	CloudBees            *map[string]interface{} `json:"cloudbees,omitempty"`
	Concur               *map[string]interface{} `json:"concur,omitempty"`
	Dropbox              *map[string]interface{} `json:"dropbox,omitempty"`
	EchoSign             *map[string]interface{} `json:"echosign,omitempty"`
	Egnyte               *map[string]interface{} `json:"egnyte,omitempty"`
	NewRelic             *map[string]interface{} `json:"newrelic,omitempty"`
	Office365            *map[string]interface{} `json:"office365,omitempty"`
	Salesforce           *map[string]interface{} `json:"salesforce,omitempty"`
	SalesforceApi        *map[string]interface{} `json:"salesforce_api,omitempty"`
	SalesforceSandboxApi *map[string]interface{} `json:"salesforce_sandbox_api,omitempty"`
	SamlP                *map[string]interface{} `json:"samlp,omitempty"`
	Layer                *map[string]interface{} `json:"layer,omitempty"`
	SapApi               *map[string]interface{} `json:"sap_api,omitempty"`
	SharePoint           *map[string]interface{} `json:"sharepoint,omitempty"`
	SpringCm             *map[string]interface{} `json:"springcm,omitempty"`
	Wams                 *map[string]interface{} `json:"wams,omitempty"`
	WsFed                *map[string]interface{} `json:"wsfed,omitempty"`
	Zendesk              *map[string]interface{} `json:"zendesk,omitempty"`
	Zoom                 *map[string]interface{} `json:"zoom,omitempty"`
}

type Addon struct{}

type AddonAws struct {
	*Addon
}

type AddonFirebase struct {
	Secret            StringOrNil `json:"secret,omitempty"`
	ClientEmail       StringOrNil `json:"client_email,omitempty"`
	PrivateKey        StringOrNil `json:"private_key,omitempty"`
	PrivateKeyId      StringOrNil `json:"private_key_id,omitempty"`
	LifetimeInSeconds Int64OrNil  `json:"lifetime_in_seconds,omitempty"`
	*Addon
}
