package vplatform

import (
	"fmt"
	"io"
	"net/http"
)

// Client ...
type Client struct {
	BaseURL               string
	VPlatformPartnerToken string
	VPlatformAuthToken    string
	httpClient            *http.Client
}

// Response - response object
type Response struct {
	Data             interface{}       `json:"data"`
	Message          string            `json:"message"`
	ValidationErrors []ValidationError `json:"validationErrors"`
	StackTrace       string            `json:"stackTrace"`
}

// ValidationError - validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Errors - holds multiple errors
type Errors []error

func (e Errors) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}

	msg := "multiple errors:"
	for _, err := range e {
		msg += "\n" + err.Error()
	}
	return msg
}

// NewClient - returns an instance of vplatform client
func NewClient(baseURL string, partnerToken string, authToken string) *Client {

	return &Client{
		BaseURL:               baseURL,
		VPlatformPartnerToken: partnerToken,
		VPlatformAuthToken:    authToken,
		httpClient:            &http.Client{},
	}
}

func makeRequest(client *Client, method string, url string, requestBody io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", client.BaseURL, url),
		requestBody,
	)
	if err != nil {
		return nil, err
	}

	request.Header.Set("vplatform-partner-token", client.VPlatformPartnerToken)
	request.Header.Set("vplatform-auth-token", client.VPlatformAuthToken)

	return client.httpClient.Do(request)
}

func makeRequestWithHeaders(
	client *Client, method string, url string, requestBody io.Reader, headers map[string]string,
) (*http.Response, error) {
	request, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", client.BaseURL, url),
		requestBody,
	)
	if err != nil {
		return nil, err
	}

	request.Header.Set("vplatform-partner-token", client.VPlatformPartnerToken)
	request.Header.Set("vplatform-auth-token", client.VPlatformAuthToken)

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	return client.httpClient.Do(request)
}
