package vplatform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Vcode ...
type Vcode struct {
	TotalScansCount  uint
	UniqueScansCount uint
	TotalOpensCount  uint
	isPublic         bool
	DateCreated      time.Time
	LastScanDate     time.Time
	Description      string
	UTI              string
	Status           string
	ActiveCount      uint
	NotActiveCount   uint
	UnassignedCount  uint
}

// CreateVcodeRequest - request body sent to create a vcode
type CreateVcodeRequest struct {
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

// Create a vcode
func (vc *Vcode) Create(client *Client, description string, quantity uint) Errors {
	if description == "" {
		return Errors{errors.New("description not found")}
	}

	requestBody, err := json.Marshal(CreateVcodeRequest{
		Description: description,
		Quantity:    quantity,
	})
	if err != nil {
		return Errors{err}
	}

	response, err := makeRequestWithHeaders(
		client,
		"POST",
		"/vcode/create",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json-patch+json",
		},
	)
	if err != nil {
		return Errors{err}
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
		vc.UTI = data
		return nil
	}
	return Errors{errors.New("response.data is not a string or it's empty")}
}
