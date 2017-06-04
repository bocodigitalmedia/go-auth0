package auth0mgmt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

type Api struct {
	sling         *sling.Sling
	authorization string
}

type Result interface{}
type Params interface{}

type SlingDecorator func(*sling.Sling) *sling.Sling

func (s *Api) Base() *sling.Sling {
	return s.sling.New().Set("Authorization", s.authorization)
}

func (s *Api) Receive(decorate SlingDecorator, result *interface{}) (*http.Response, error) {
	apiErr := new(ApiError)
	base := decorate(s.Base())
	resp, err := base.Receive(result, apiErr)

	if err == nil && !apiErr.IsEmpty() {
		err = apiErr
	}

	return resp, err
}

func (s *Api) Get(path string, params, result interface{}) (*http.Response, error) {
	decorate := func(s *sling.Sling) *sling.Sling {
		return s.Get(path).QueryStruct(params)
	}

	return s.Receive(decorate, &result)
}

func (s *Api) Post(path string, params, result interface{}) (*http.Response, error) {
	decorate := func(s *sling.Sling) *sling.Sling {
		return s.Post(path).BodyJSON(&params)
	}

	return s.Receive(decorate, &result)
}

func (s *Api) Patch(path string, params, result interface{}) (*http.Response, error) {
	decorate := func(s *sling.Sling) *sling.Sling {
		return s.Patch(path).BodyJSON(&params)
	}

	return s.Receive(decorate, &result)
}

func (s *Api) Delete(path string, params, result interface{}) (*http.Response, error) {
	decorate := func(s *sling.Sling) *sling.Sling {
		return s.Delete(path).QueryStruct(&params)
	}

	return s.Receive(decorate, &result)
}

type NewApiParams struct {
	Domain      string
	AccessToken string
	HttpClient  *http.Client
}

func NewApi(params *NewApiParams) *Api {
	baseUrl := fmt.Sprintf("https://%s/api/v2/", params.Domain)
	authorization := fmt.Sprintf("Bearer %s", params.AccessToken)
	httpClient := params.HttpClient

	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Second * 60}
	}

	return &Api{
		sling:         sling.New().Base(baseUrl).Client(httpClient),
		authorization: authorization,
	}
}
