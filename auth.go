package vplatform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// AuthService - auth service
type AuthService struct {
	client *Client
}

// LoginInputModel ...
type LoginInputModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login - login requests user token given username and password
func (aus *AuthService) Login(request *LoginInputModel) (*Result, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := aus.client.makeRequestWithHeaders(
		"POST",
		"/auth/login",
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

	if data, ok := responseBody.Data.(string); ok && data != "" {
		return responseBody, nil
	}
	return nil, Errors{errors.New("response.data is not a string or it's empty")}
}

// Logout - logout requests user token given username and password
func (aus *AuthService) Logout(request *LoginInputModel) (*Result, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := aus.client.makeRequestWithHeaders(
		"POST",
		"/auth/login",
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

	if data, ok := responseBody.Data.(string); ok && data != "" {
		return responseBody, nil
	}
	return nil, Errors{errors.New("response.data is not a string or it's empty")}
}
