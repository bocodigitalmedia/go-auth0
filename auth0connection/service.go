package auth0connection

import (
	"net/http"
	"strings"

	"github.com/bocodigitalmedia/go-auth0/auth0mgmt"
)

const ServicePath = "connections"

type Service struct {
	Api *auth0mgmt.Api
}

func (s *Service) Path(paths ...string) string {
	strs := append([]string{ServicePath}, paths...)
	return strings.Join(strs, "/")
}

func (s *Service) List(params *ListParams) (*[]Connection, *http.Response, error) {
	path := s.Path()
	result := new([]Connection)

	resp, err := s.Api.Get(path, params, result)
	return result, resp, err
}

func (s *Service) Create(params *Properties) (*Connection, *http.Response, error) {
	path := s.Path()
	result := new(Connection)
	resp, err := s.Api.Post(path, params, result)
	return result, resp, err
}

func (s *Service) Read(id string, params *ReadParams) (*Connection, *http.Response, error) {
	path := s.Path(id)
	result := new(Connection)
	resp, err := s.Api.Get(path, params, result)
	return result, resp, err
}

func (s *Service) Update(id string, params *Properties) (*Connection, *http.Response, error) {
	path := s.Path(id)
	result := new(Connection)
	resp, err := s.Api.Patch(path, params, result)
	return result, resp, err
}

func (s *Service) Delete(id string) (*http.Response, error) {
	path := s.Path(id)
	resp, err := s.Api.Delete(path, nil, nil)
	return resp, err
}

func (s *Service) DeleteUser(id string, params *DeleteUserParams) (*http.Response, error) {
	path := s.Path(id, "users")
	resp, err := s.Api.Delete(path, params, nil)
	return resp, err
}

type ListParams struct {
	PerPage       interface{} `url:"per_page,omitempty"`
	Page          interface{} `url:"page,omitempty"`
	Strategy      *[]string   `url:"strategy,omitempty,comma"`
	Fields        interface{} `url:"fields,omitempty"`
	IncludeFields interface{} `url:"include_fields,omitempty"`
}

type ReadParams struct {
	Fields        interface{} `url:"fields,omitempty"`
	IncludeFields interface{} `url:"include_fields,omitempty"`
}

type DeleteUserParams struct {
	Email string `url:"email"`
}
