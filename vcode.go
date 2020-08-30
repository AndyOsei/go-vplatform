package vplatform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// VCodeService - vcode service
type VCodeService struct {
	client *Client
}

// VCodeCreateInputModel - VCode creation model
type VCodeCreateInputModel struct {
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

// VCodeOutputModel - VCode information and assigned packages.
type VCodeOutputModel struct {
	Description string               `json:"description"`
	Packages    []PackageOutputModel `json:"packages"`
	UTI         string               `json:"uti"`
}

// VCodeDetailsOutputModel - VCode details
type VCodeDetailsOutputModel struct {
	TotalScansCount  uint      `json:"totalScansCount"`
	UniqueScansCount uint      `json:"uniqueScansCount"`
	TotalOpensCount  uint      `json:"totalOpensCount"`
	IsPublic         bool      `json:"isPublic"`
	DateCreated      time.Time `json:"dateCreated"`
	LastScanDate     time.Time `json:"lastScanDate"`
	Description      string    `json:"description"`
	UTI              string    `json:"uti"`
	Status           string    `json:"status"`
	ActiveCount      uint      `json:"activeCount"`
	NotActiveCount   uint      `json:"notActiveCount"`
	UnassignedCount  uint      `json:"unassignedCount"`
}

// Create - create a vcode
func (vcs *VCodeService) Create(request *VCodeCreateInputModel) (*Result, error) {

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, Errors{err}
	}

	response, err := vcs.client.makeRequestWithHeaders(
		"POST",
		"/vcode/create",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json-patch+json",
		},
	)
	if err != nil {
		return nil, Errors{err}
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

// GetDetailsByUTI - Get VCode details by uti
func (vcs *VCodeService) GetDetailsByUTI(uti string) (*Result, error) {
	response, err := vcs.client.makeRequestWithURLParams(
		"GET",
		fmt.Sprintf("/vcode/%s", uti),
	)
	if err != nil {
		return nil, Errors{err}
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

	if _, ok := responseBody.Data.(VCodeDetailsOutputModel); ok {
		return responseBody, nil
	}
	return nil, Errors{errors.New("response.data is not of type Vcode")}
}
