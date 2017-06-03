package auth0mgmt

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type Service struct {
	sling *sling.Sling
	token string
}

type Result interface{}
type Params interface{}

type SlingDecorator func(*sling.Sling) *sling.Sling

func (s *Service) Base() *sling.Sling {
	authorization := fmt.Sprintf("Bearer %s", s.token)
	return s.sling.New().Set("Authorization", authorization)
}

func (s *Service) Receive(decorate SlingDecorator, result *interface{}) (*http.Response, error) {
	svcerr := new(ServiceError)
	base := decorate(s.Base())
	resp, err := base.Receive(result, svcerr)

	if err != nil && !svcerr.IsEmpty() {
		err = svcerr
	}

	return resp, err
}

func (s *Service) Get(path string, params, result interface{}) (*http.Response, error) {
	decorate := func(s *sling.Sling) *sling.Sling {
		return s.Get(path).QueryStruct(params)
	}

	return s.Receive(decorate, &result)
}

func (s *Service) Post(path string, params, result interface{}) (*http.Response, error) {
	decorate := func(s *sling.Sling) *sling.Sling {
		return s.Post(path).BodyJSON(&params)
	}

	return s.Receive(decorate, &result)
}

func (s *Service) Patch(path string, params, result interface{}) (*http.Response, error) {
	decorate := func(s *sling.Sling) *sling.Sling {
		return s.Patch(path).BodyJSON(&params)
	}

	return s.Receive(decorate, &result)
}

func (s *Service) Delete(path string, params, result interface{}) (*http.Response, error) {
	decorate := func(s *sling.Sling) *sling.Sling {
		return s.Delete(path).QueryStruct(&params)
	}

	return s.Receive(decorate, &result)
}

func NewService(httpClient *http.Client, domain, token string) *Service {
	baseUrl := fmt.Sprintf("https://%s/api/v2/", domain)

	return &Service{
		sling: sling.New().Base(baseUrl),
		token: token,
	}
}
