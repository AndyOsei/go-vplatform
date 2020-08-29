package vplatform

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"bytes"
)

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

// Create a package
func (pkg *Package) Create(client *Client, request *CreatePackageRequest) error {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	response, err := makeRequestWithHeaders(
		client,
		"POST",
		"/package/create",
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

	if data, ok := responseBody.Data.(string); ok && data != "" {
		pkg.PublicID = data
		return nil
	}
	return Errors{errors.New("response.data is not a string or it's empty")}

}
