package vplatform

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"bytes"
)

// PackageService - package service
type PackageService struct {
	client *Client
}

// Package - package object
type Package struct {
	PublicID string `json:"publicId"`
}

// CreatePackageRequest ...
type CreatePackageRequest struct {
	PublicID       string `json:"publicId"`
	Description    string `json:"description"`
	UTI            string `json:"uti"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	ActionPublicID string `json:"actionPublicId"`
	TimeRules      []struct {
		publicID string
	} `json:"timeRules"`
	LocationRules []struct {
		publicID string
	} `json:"locationRules"`
	RestrictedEmails string    `json:"restrictedEmails"`
	Platform         string    `json:"platform"`
	SingleScan       string    `json:"singleScan"`
	BillingStartDate time.Time `json:"billingStartDate"`
	BillingEndDate   time.Time `json:"billingEndDate"`
	MaxumberOfScans  uint      `json:"maxNumberOfScans"`
}

// PackageOutputModel ...
type PackageOutputModel struct {
	PublicID string      `json:"publicId"`
	Title    string      `json:"title"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
}

// Create - create a package
func (pks *PackageService) Create(request *CreatePackageRequest) (*Result, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := pks.client.makeRequestWithHeaders(
		"POST",
		"/package/create",
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

	if data, ok := responseBody.Data.(string); ok && data != "" {
		return responseBody, nil
	}

	return nil, Errors{errors.New("response.data is not a string or it's empty")}

}
