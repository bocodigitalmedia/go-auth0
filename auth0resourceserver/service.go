package auth0resourceserver

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
)

const ServicePath = "resource-servers"

type Service struct {
	Api *auth0mgmt.Api
}

func (s *Service) Path(paths ...string) string {
	strs := append([]string{ServicePath}, paths...)
	return strings.Join(strs, "/")
}

func (s *Service) List() (*[]ResourceServer, *http.Response, error) {
	path := s.Path()
	result := new([]ResourceServer)
	resp, err := s.Api.Get(path, nil, result)
	return result, resp, err
}

func (s *Service) Create(params *Properties) (*ResourceServer, *http.Response, error) {
	path := s.Path()
	result := new(ResourceServer)
	resp, err := s.Api.Post(path, params, result)
	return result, resp, err
}

func (s *Service) Read(idOrAudience string) (*ResourceServer, *http.Response, error) {
	path := s.Path(url.PathEscape(idOrAudience))
	result := new(ResourceServer)
	resp, err := s.Api.Get(path, nil, result)
	return result, resp, err
}

func (s *Service) Update(idOrAudience string, params *Properties) (*ResourceServer, *http.Response, error) {
	path := s.Path(url.PathEscape(idOrAudience))
	result := new(ResourceServer)
	resp, err := s.Api.Patch(path, params, result)
	return result, resp, err
}

func (s *Service) Delete(idOrAudience string) (*http.Response, error) {
	path := s.Path(url.PathEscape(idOrAudience))
	resp, err := s.Api.Delete(path, nil, nil)
	return resp, err
}
