package auth0tenantsettings

import (
	"net/http"
	"strings"

	"github.com/bocodigitalmedia/auth0/auth0mgmt"
)

const ServicePath = "tenants/settings"

type Service struct {
	Api *auth0mgmt.Api
}

func (s *Service) Path(paths ...string) string {
	strs := append([]string{ServicePath}, paths...)
	return strings.Join(strs, "/")
}

func (s *Service) Read(params *ReadParams) (*Settings, *http.Response, error) {
	path := s.Path()
	result := new(Settings)
	resp, err := s.Api.Get(path, params, result)
	return result, resp, err
}

func (s *Service) Update(params *Settings) (*Settings, *http.Response, error) {
	path := s.Path()
	result := new(Settings)
	resp, err := s.Api.Patch(path, params, result)
	return result, resp, err
}

type ReadParams struct {
	Fields        interface{} `url:"fields,omitempty"`
	IncludeFields interface{} `url:"include_fields,omitempty"`
}
