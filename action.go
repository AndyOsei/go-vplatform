package vplatform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// Action ...
type Action struct {
	PublicID               string                 `json:"publicId"`
	Description            string                 `json:"description"`
	Type                   string                 `json:"type"`
	Data                   map[string]interface{} `json:"data"`
	PackageUsage           uint                   `json:"packageUsage"`
	UnassignedPackageUsage uint                   `json:"unassignedPackageUsage"`
}

// URLActionRequest - url action request body
type URLActionRequest struct {
	URL          string `json:"url"`
	InAppBrowser bool   `json:"inapp_browser"`
	PublicID     string `json:"publicId"`
	Description  string `json:"description"`
}

// AppLinkActionRequest - app link action request body
type AppLinkActionRequest struct {
	AppLink     string `json:"appLink"`
	IosLink     string `json:"iosLink"`
	AndroidLink string `json:"androidLink"`
	PublicID    string `json:"publicId"`
	Description string `json:"description"`
}

// ContactActionRequest - contact card action request body
type ContactActionRequest struct {
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
func (at *Action) CreateURLAction(client *Client, request *URLActionRequest) error {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	response, err := makeRequestWithHeaders(
		client,
		"POST",
		"/action/url",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json-patch+json",
		},
	)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	responseBody := new(Response)
	json.NewDecoder(response.Body).Decode(responseBody)

	if len(responseBody.ValidationErrors) > 0 {
		errors := Errors{}
		for _, err := range responseBody.ValidationErrors {
			errors = append(errors, fmt.Errorf("%s: %s", err.Field, err.Message))
		}
		return errors
	}

	if data, ok := responseBody.Data.(Action); ok {
		at = &data
		return nil
	}
	return Errors{errors.New("response.data is not of type Action")}
}

// CreateAppLinkAction - create app link action
func (at *Action) CreateAppLinkAction(client *Client, request *AppLinkActionRequest) error {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	response, err := makeRequestWithHeaders(
		client,
		"POST",
		"/action/applink",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json-patch+json",
		},
	)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	responseBody := new(Response)
	json.NewDecoder(response.Body).Decode(responseBody)

	if len(responseBody.ValidationErrors) > 0 {
		errors := Errors{}
		for _, err := range responseBody.ValidationErrors {
			errors = append(errors, fmt.Errorf("%s: %s", err.Field, err.Message))
		}
		return errors
	}

	if data, ok := responseBody.Data.(Action); ok {
		at = &data
		return nil
	}
	return Errors{errors.New("response.data is not of type Action")}
}

// CreateContactAction - create contact card action
func (at *Action) CreateContactAction(client *Client, request *ContactActionRequest) error {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	response, err := makeRequestWithHeaders(
		client,
		"POST",
		"/action/contact",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json-patch+json",
		},
	)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	responseBody := new(Response)
	json.NewDecoder(response.Body).Decode(responseBody)

	if len(responseBody.ValidationErrors) > 0 {
		errors := Errors{}
		for _, err := range responseBody.ValidationErrors {
			errors = append(errors, fmt.Errorf("%s: %s", err.Field, err.Message))
		}
		return errors
	}

	if data, ok := responseBody.Data.(Action); ok {
		at = &data
		return nil
	}
	return Errors{errors.New("response.data is not of type Action")}
}
