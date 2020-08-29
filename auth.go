package vplatform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// Auth ...
type Auth struct{}

// LoginInputModel ...
type LoginInputModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login - login requests user token given username and password
func (tr *Auth) Login(client *Client, request *LoginInputModel) error {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	response, err := makeRequestWithHeaders(
		client,
		"POST",
		"/auth/login",
		bytes.NewBuffer(requestBody),
		map[string]string{
			"Content-Type": "application/json",
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

	if data, ok := responseBody.Data.(Auth); ok {
		tr = &data
		return nil
	}
	return Errors{errors.New("response.data is not of type LoginInputModel")}
}
