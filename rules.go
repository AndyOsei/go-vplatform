package vplatform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// RuleService - rule service
type RuleService struct {
	client *Client
}

// TimeRule ...
type TimeRule struct {
	StartDate              string    `json:"startDate"`
	ExpiryDate             string    `json:"expiryDate"`
	FromHour               uint      `json:"fromHour"`
	FromMinutes            uint      `json:"fromMinutes"`
	ToHour                 uint      `json:"toHour"`
	ToMinutes              uint      `json:"toMinutes"`
	PublicID               string    `json:"publicId"`
	Description            string    `json:"description"`
	DateCreated            time.Time `json:"dateCreated"`
	DateUpdated            time.Time `json:"dateUpdated"`
	PackageUsage           uint      `json:"packageUsage"`
	UnassignedPackageUsage uint      `json:"unassignedPackageUsage"`
}

// TimeRuleRequest - request body for creating a time rule
type TimeRuleRequest struct {
	StartDate   string `json:"startDate"`
	ExpiryDate  string `json:"expiryDate"`
	FromHour    uint   `json:"fromHour"`
	FromMinutes uint   `json:"fromMinutes"`
	ToHour      uint   `json:"toHour"`
	ToMinutes   uint   `json:"toMinutes"`
	Description string `json:"description"`
}

//GeoFenceRule - Geofence rule object
type GeoFenceRule struct {
	Type        string `json:"type"`
	RadiusModel struct {
		Location struct {
			Lat int `json:"lat"`
			Lng int `json:"lng"`
		} `json:"location"`
		Radius int `json:"radius"`
	} `json:"radiusModel"`
	PolygonModel struct {
		Points []struct {
			Lat int `json:"lat"`
			Lng int `json:"lng"`
		} `json:"points"`
	} `json:"polygonModel"`
	PublicID               string    `json:"publicId"`
	Description            string    `json:"description"`
	DateCreated            time.Time `json:"dateCreated"`
	DateUpdated            time.Time `json:"dateUpdated"`
	PackageUsage           uint      `json:"packageUsage"`
	UnassignedPackageUsage uint      `json:"unassignedPackageUsage"`
}

//GeoFenceRuleRequest - request body for creating a geofence rule
type GeoFenceRuleRequest struct {
	Type        string `json:"type"`
	RadiusModel struct {
		Location struct {
			Lat int `json:"lat"`
			Lng int `json:"lng"`
		} `json:"location"`
		Radius int `json:"radius"`
	} `json:"radiusModel"`
	PolygonModel struct {
		Points []struct {
			Lat int `json:"lat"`
			Lng int `json:"lng"`
		} `json:"points"`
	} `json:"polygonModel"`
	Description string `json:"description"`
}

// CreateTimeRule - create a time rule
func (rls *RuleService) CreateTimeRule(request *TimeRuleRequest) (*Result, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := rls.client.makeRequestWithHeaders(
		"POST",
		"/rules/time/create",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json",
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

	if _, ok := responseBody.Data.(TimeRule); ok {
		return responseBody, nil
	}
	return nil, Errors{errors.New("response.data is not of type TimeRule")}
}

// Create a geofence rule
func (rls *RuleService) Create(request *GeoFenceRuleRequest) (*Result, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := rls.client.makeRequestWithHeaders(
		"POST",
		"/rules/time/create",
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

	if _, ok := responseBody.Data.(GeoFenceRule); ok {
		return responseBody, nil
	}
	return nil, Errors{errors.New("response.data is not of type GeoFenceRule")}
}
