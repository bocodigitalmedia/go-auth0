package auth0mgmt

import "strings"

type Service struct {
	Api      *Api
	BasePath string
}

func (s *Service) Path(paths ...string) string {
	strs := append([]string{s.BasePath}, paths...)
	return strings.Join(strs, "/")
}

type NewServiceParams struct {
	Api *Api
}

type ListParams struct {
	Page          interface{} `url:"page,omitempty"`
	PerPage       interface{} `url:"per_page,omitempty"`
	IncludeTotals interface{} `url:"include_totals,omitempty"`
	ReadParams
}

type ReadParams struct {
	Fields        interface{} `url:"fields,omitempty"`
	IncludeFields interface{} `url:"include_fields,omitempty"`
}
