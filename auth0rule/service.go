package auth0rule

import (
	"net/http"
	"strings"

	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
)

const ServicePath = "rules"

type Service struct {
	Api *auth0mgmt.Api
}

func (s *Service) Path(paths ...string) string {
	strs := append([]string{ServicePath}, paths...)
	return strings.Join(strs, "/")
}

func (s *Service) List(params *ListParams) (*[]Rule, *http.Response, error) {
	path := s.Path()
	data := new([]Rule)
	resp, err := s.Api.Get(path, &params, data)

	return data, resp, err
}

func (s *Service) Create(params *Properties) (*Rule, *http.Response, error) {
	path := s.Path()
	data := new(Rule)
	resp, err := s.Api.Post(path, params, data)
	return data, resp, err
}

func (s *Service) Read(id string, params *ReadParams) (*Rule, *http.Response, error) {
	path := s.Path(id)
	data := new(Rule)
	resp, err := s.Api.Get(path, params, data)
	return data, resp, err
}

func (s *Service) Update(id string, params *Properties) (*Rule, *http.Response, error) {
	path := s.Path(id)
	data := new(Rule)
	resp, err := s.Api.Patch(path, params, data)
	return data, resp, err
}

func (s *Service) Delete(id string) (*http.Response, error) {
	path := s.Path(id)
	resp, err := s.Api.Delete(path, nil, nil)
	return resp, err
}

type ListParams struct {
	Enabled       interface{} `json:"enabled,omitempty"`
	Fields        interface{} `json:"fields,omitempty"`
	IncludeFields interface{} `json:"include_fields,omitempty"`
}

type ReadParams struct {
	Fields        interface{} `json:"fields,omitempty"`
	IncludeFields interface{} `json:"include_fields,omitempty"`
}
