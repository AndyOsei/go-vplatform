package vplatform

import (
	"fmt"
	"io"
	"net/http"
)

// Client - vplatform client object
type Client struct {
	httpClient *http.Client

	BaseURL               string
	VPlatformPartnerToken string
	VPlatformAuthToken    string

	Auth    *AuthService
	VCode   *VCodeService
	Package *PackageService
	Action  *ActionService
	Rule    *RuleService
}

// Result - Response Model
type Result struct {
	Data             interface{}              `json:"data"`
	Message          string                   `json:"message"`
	ValidationErrors []ValidationErrorMessage `json:"validationErrors"`
	StackTrace       string                   `json:"stackTrace"`
}

// ValidationErrorMessage - validation error message model
type ValidationErrorMessage struct {
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

	client := &Client{
		httpClient:            &http.Client{},
		BaseURL:               baseURL,
		VPlatformPartnerToken: partnerToken,
		VPlatformAuthToken:    authToken,
	}

	client.Auth = &AuthService{client: client}
	client.Package = &PackageService{client: client}
	client.VCode = &VCodeService{client: client}
	client.Action = &ActionService{client: client}
	client.Rule = &RuleService{client: client}

	return client
}

func (cl *Client) makeRequest(method string, url string, requestBody io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", cl.BaseURL, url),
		requestBody,
	)
	if err != nil {
		return nil, err
	}

	request.Header.Set("vplatform-partner-token", cl.VPlatformPartnerToken)
	request.Header.Set("vplatform-auth-token", cl.VPlatformAuthToken)

	return cl.httpClient.Do(request)
}

func (cl *Client) makeRequestWithHeaders(
	method string, url string, requestBody io.Reader, headers map[string]string,
) (*http.Response, error) {
	request, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", cl.BaseURL, url),
		requestBody,
	)
	if err != nil {
		return nil, err
	}

	request.Header.Set("vplatform-partner-token", cl.VPlatformPartnerToken)
	request.Header.Set("vplatform-auth-token", cl.VPlatformAuthToken)

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	return cl.httpClient.Do(request)
}

func (cl *Client) makeRequestWithURLParams(method string, url string) (*http.Response, error) {
	request, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", cl.BaseURL, url),
		nil,
	)
	if err != nil {
		return nil, err
	}

	request.Header.Set("vplatform-partner-token", cl.VPlatformPartnerToken)
	request.Header.Set("vplatform-auth-token", cl.VPlatformAuthToken)

	return cl.httpClient.Do(request)
}
