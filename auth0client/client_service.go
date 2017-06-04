package auth0client

import (
	"fmt"
	"net/http"

	"github.com/bocodigitalmedia/auth0/auth0mgmt"
)

type Service struct {
	api      *auth0mgmt.Api
	basePath string
}

func (s *Service) List(params *ListParams) (*[]Client, *http.Response, error) {
	path := s.basePath
	result := new([]Client)
	resp, err := s.api.Get(path, params, result)
	return result, resp, err
}

func (s *Service) Create(params *Properties) (*Client, *http.Response, error) {
	path := s.basePath
	result := new(Client)
	resp, err := s.api.Post(path, params, result)
	return result, resp, err
}

func (s *Service) Read(id string, params *ReadParams) (*Client, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, id)
	result := new(Client)
	resp, err := s.api.Get(path, params, result)
	return result, resp, err
}

func (s *Service) Update(id string, params *Properties) (*Client, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, id)
	result := new(Client)
	resp, err := s.api.Patch(path, params, result)
	return result, resp, err
}

func (s *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, id)
	resp, err := s.api.Delete(path, nil, nil)
	return resp, err
}

func (s *Service) RotateSecret(id string) (*Client, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/rotate-secret", s.basePath, id)
	result := new(Client)
	resp, err := s.api.Post(path, nil, result)
	return result, resp, err
}

const ServiceBasePath = "clients"

type NewServiceParams struct {
	Api *auth0mgmt.Api
}

func NewService(params *NewServiceParams) *Service {
	return &Service{
		api:      params.Api,
		basePath: ServiceBasePath,
	}
}

type ListParams struct {
	Page          Int64OrNil `json:"page,omitempty"`
	PerPage       Int64OrNil `json:"per_page,omitempty"`
	IncludeTotals BoolOrNil  `json:"include_totals,omitempty"`
	*ReadParams
}

type ReadParams struct {
	Fields        StringOrNil `json:"fields,omitempty"`
	IncludeFields BoolOrNil   `json:"include_fields,omitempty"`
}
