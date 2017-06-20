package auth0client

import (
	"net/http"
	"strings"

	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
)

const ServicePath = "clients"

type Service struct {
	Api *auth0mgmt.Api
}

func (s *Service) Path(paths ...string) string {
	strs := append([]string{ServicePath}, paths...)
	return strings.Join(strs, "/")
}

func (s *Service) List(params *ListParams) (*[]Client, *http.Response, error) {
	path := s.Path()
	result := new([]Client)
	resp, err := s.Api.Get(path, params, result)

	return result, resp, err
}

func (s *Service) Create(params *Properties) (*Client, *http.Response, error) {
	path := s.Path()
	result := new(Client)
	resp, err := s.Api.Post(path, params, result)
	return result, resp, err
}

func (s *Service) Read(id string, params *ReadParams) (*Client, *http.Response, error) {
	path := s.Path(id)
	result := new(Client)
	resp, err := s.Api.Get(path, params, result)
	return result, resp, err
}

func (s *Service) Update(id string, params *Properties) (*Client, *http.Response, error) {
	path := s.Path(id)
	result := new(Client)
	resp, err := s.Api.Patch(path, params, result)
	return result, resp, err
}

func (s *Service) Delete(id string) (*http.Response, error) {
	path := s.Path(id)
	resp, err := s.Api.Delete(path, nil, nil)
	return resp, err
}

func (s *Service) RotateSecret(id string) (*Client, *http.Response, error) {
	path := s.Path(id, "rotate-secret")
	result := new(Client)
	resp, err := s.Api.Post(path, nil, result)
	return result, resp, err
}

type ListParams struct {
	Page          int64       `url:"page,omitempty"`
	PerPage       int64       `url:"per_page,omitempty"`
	Fields        interface{} `url:"fields,omitempty"`
	IncludeFields interface{} `url:"include_fields,omitempty"`
}

type ReadParams struct {
	Fields        interface{} `url:"fields,omitempty"`
	IncludeFields interface{} `url:"include_fields,omitempty"`
}
