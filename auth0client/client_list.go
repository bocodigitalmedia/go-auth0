package auth0client

type ClientList struct {
	Total   int64     `json:"total"`
	Page    int64     `json:"start"`
	PerPage int64     `json:"limit"`
	Clients *[]Client `json:"clients"`
}
