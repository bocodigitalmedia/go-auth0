package auth0mgmt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

type Api struct {
	sling *sling.Sling
	token string
}

type Result interface{}
type Params interface{}

type SlingDecorator func(*sling.Sling) *sling.Sling

func (s *Api) Base() *sling.Sling {
	authorization := fmt.Sprintf("Bearer %s", s.token)
	return s.sling.New().Set("Authorization", authorization)
}

func (s *Api) Receive(decorate SlingDecorator, result *interface{}) (*http.Response, error) {
	apiErr := new(ApiError)
	base := decorate(s.Base())
	resp, err := base.Receive(result, apiErr)

	if err == nil && !apiErr.IsEmpty() {
		b, _ := json.MarshalIndent(apiErr, "", "\t")

		log.Println(apiErr)
		log.Println(string(b))
		log.Printf("[ERROR] %s", apiErr.Error())
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
	httpClient := params.HttpClient

	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Second * 60}
	}

	return &Api{
		sling: sling.New().Base(baseUrl).Client(httpClient),
		token: params.AccessToken,
	}
}
