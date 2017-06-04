package auth0tenant

import (
	"net/http"
	"strings"

	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
)

const ServicePath = "tenants/settings"

type Service struct {
	Api *auth0mgmt.Api
}

func (s *Service) Path(paths ...string) string {
	strs := append([]string{ServicePath}, paths...)
	return strings.Join(strs, "/")
}

func (s *Service) ReadSettings(params *ReadSettingsParams) (*Settings, *http.Response, error) {
	path := s.Path()
	result := new(Settings)
	resp, err := s.Api.Get(path, params, result)
	return result, resp, err
}

func (s *Service) UpdateSettings(params *Settings) (*Settings, *http.Response, error) {
	path := s.Path()
	result := new(Settings)
	resp, err := s.Api.Patch(path, params, result)
	return result, resp, err
}

type ReadSettingsParams struct {
	Fields        interface{} `url:"fields,omitempty"`
	IncludeFields interface{} `url:"include_fields,omitempty"`
}
