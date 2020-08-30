package vplatform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// ActionService - action service
type ActionService struct {
	client *Client
}

// ActionOutputModel ...
type ActionOutputModel struct {
	PublicID               string      `json:"publicId"`
	Description            string      `json:"description"`
	Type                   string      `json:"type"`
	Data                   interface{} `json:"data"`
	PackageUsage           uint        `json:"packageUsage"`
	UnassignedPackageUsage uint        `json:"unassignedPackageUsage"`
}

// URLActionDataInputModel - url action input model
type URLActionDataInputModel struct {
	URL          string `json:"url"`
	InAppBrowser bool   `json:"inapp_browser"`
	PublicID     string `json:"publicId"`
	Description  string `json:"description"`
} // UrlActionDataInputModel

// AppLinkActionDataInputModel - app link action input model
type AppLinkActionDataInputModel struct {
	AppLink     string `json:"appLink"`
	IosLink     string `json:"iosLink"`
	AndroidLink string `json:"androidLink"`
	PublicID    string `json:"publicId"`
	Description string `json:"description"`
}

// ContactActionDataInputModel - contact card action input model
type ContactActionDataInputModel struct {
	Name                 string `json:"name"`
	Title                string `json:"title"`
	Phone                string `json:"phone"`
	Mobile               string `json:"mobile"`
	Email                string `json:"email"`
	Organization         string `json:"organisation"`
	Position             string `json:"position"`
	Address              string `json:"address"`
	Website              string `json:"website"`
	Image                string `json:"image"`
	Lat                  int    `json:"lat"`
	Lng                  int    `json:"lng"`
	ImageProfilePublicID string `json:"imageProfilePublicId"`
	PublicID             string `json:"publicId"`
	Description          string `json:"description"`
}

// CreateURLAction - create url action
func (ats *ActionService) CreateURLAction(request *URLActionDataInputModel) (*Result, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := ats.client.makeRequestWithHeaders(
		"POST",
		"/action/url",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json-patch+json",
		},
	)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody := new(Result)
	json.NewDecoder(response.Body).Decode(responseBody)

	if len(responseBody.ValidationErrors) > 0 {
		errors := Errors{}
		for _, err := range responseBody.ValidationErrors {
			errors = append(errors, fmt.Errorf("%s: %s", err.Field, err.Message))
		}
		return nil, errors
	}

	if _, ok := responseBody.Data.(ActionOutputModel); ok {
		return responseBody, nil
	}
	return nil, Errors{errors.New("response.data is not of type Action")}
}

// CreateAppLinkAction - create app link action
func (ats *ActionService) CreateAppLinkAction(request *AppLinkActionDataInputModel) (*Result, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := ats.client.makeRequestWithHeaders(
		"POST",
		"/action/applink",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json-patch+json",
		},
	)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody := new(Result)
	json.NewDecoder(response.Body).Decode(responseBody)

	if len(responseBody.ValidationErrors) > 0 {
		errors := Errors{}
		for _, err := range responseBody.ValidationErrors {
			errors = append(errors, fmt.Errorf("%s: %s", err.Field, err.Message))
		}
		return nil, errors
	}

	if _, ok := responseBody.Data.(ActionOutputModel); ok {
		return responseBody, nil
	}
	return nil, Errors{errors.New("response.data is not of type Action")}
}

// CreateContactAction - create contact card action
func (ats *ActionService) CreateContactAction(request *ContactActionDataInputModel) (*Result, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := ats.client.makeRequestWithHeaders(
		"POST",
		"/action/contact",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json-patch+json",
		},
	)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody := new(Result)
	json.NewDecoder(response.Body).Decode(responseBody)

	if len(responseBody.ValidationErrors) > 0 {
		errors := Errors{}
		for _, err := range responseBody.ValidationErrors {
			errors = append(errors, fmt.Errorf("%s: %s", err.Field, err.Message))
		}
		return nil, errors
	}

	if _, ok := responseBody.Data.(ActionOutputModel); ok {
		return responseBody, nil
	}
	return nil, Errors{errors.New("response.data is not of type Action")}
}
