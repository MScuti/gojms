package gojms

import (
	"bytes"
	"fmt"
	"github.com/MScuti/gojms/pkg/accouts"
	"github.com/MScuti/gojms/pkg/assets"
	"github.com/MScuti/gojms/pkg/terminal"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"net/url"
)

// JmsAPIConfig represents the configuration for the JMS API.
// It contains the information about the endpoints and the authentication token.
type JmsAPIConfig struct {
	Endpoints string `json:"endpoints"`
	Token     string `json:"token"`
}

// MakeRequest creates an HTTP request with a specified method, endpoint, and data.
//
// Parameters:
//   - method: A string that represents the HTTP method (GET, POST, PUT, etc.).
//   - endpoint: A string that represents the URL of the endpoint the request is to be sent to.
//   - data: The interface that should be sent as the request body. It is marshalled using the sonic.Marshal function. If the parameter is nil, the request body will be set as nil.
//
// Returns:
//   - A pointer to the built http.Request.
//   - An error which will be non-nil in case of any errors occurring during the creation of the new http request or during the marshalling of the provided data.
//
// Implementation:
//
//	It marshals 'data' into a byte array using sonic.Marshal if 'data' is not nil.
//	Then creates a new http.Request with the provided 'method' and 'endpoint', and with the marshalled data as the body.
//	If any error occurs during these operations, it will return immediately with the respective error.
//	If 'data' is nil, it will proceed to create the new http request with a nil body.
//	Finally, before returning, it will set "Content-Type" and "Authorization" headers on the created http.Request.
func (j *JmsAPIConfig) MakeRequest(method, endpoint string, data interface{}) (*http.Request, error) {
	var err error
	var body = make([]byte, 0)

	// marshal body if data is not nil
	if data != nil {
		body, err = sonic.Marshal(data)
		if err != nil {
			return nil, err
		}
	} else {
		body = nil
	}

	// make request
	req, err := http.NewRequest(method, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// set request header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", j.Token))

	return req, nil

}

// DoRequest performs an HTTP request using the provided http.Request object.
// The function unmarshals the response body into the provided result object.
//
// Parameters:
//   - req: A pointer to the http.Request object that contains all the necessary information for the request.
//   - result: The interface where the response body will be unmarshaled into. If this parameter is nil, response body will not be unmarshalled and the function will return immediately after checking the response status code.
//
// Returns:
//
//	An error which will be non-nil in case of any errors occurred during executing the HTTP request or unmarshalling the response. If the response HTTP status code is not in the range of 200-399, it will return an error stating the response code and body content.
//
// Implementation:
//
//	The function performs the request using a new http.Client object.
//	It then reads the response body and checks the status code.
//	If the code is not in the range of 200-399, an error will be returned including the response code and body content.
//	If the result parameter is not nil, the function will attempt to unmarshal the response body into it using the sonic.Unmarshal function.
//	The function returns an error from the unmarshal operation if occurred - or nil if the operation was successful.
func (j *JmsAPIConfig) DoRequest(req *http.Request, result interface{}) error {
	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check response status code
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("server response code is not ok, code:%d, content:%s", resp.StatusCode, body)
	}

	// check if result is nil
	if result == nil {
		return nil
	}

	// unmarshal response body
	err = sonic.Unmarshal(body, result)
	return err

}

// SetQuery is a method on the JmsAPIConfig struct.
// It receives an http.Request and a set of url.Values as parameters.
// The method sets the URL query string of the given http.Request based
// on the encoded form of the received url.Values and returns the http.Request.
func (j *JmsAPIConfig) SetQuery(req *http.Request, v url.Values) *http.Request {
	// set query
	req.URL.RawQuery = v.Encode()
	return req
}

// The Terminal struct holds the Sessions object for terminal operations.
// It is used to manage and interact with terminal sessions.
type Terminal struct {
	Session terminal.Sessions
}

// The Account struct holds the Account object for account operations.
// It is used to manage and interact with accounts.
type Account struct {
	Account accouts.Account
}

// The Assets struct holds the Assets object for asset operations.
// It is used to manage and interact with assets.
type Assets struct {
	Assets assets.Assets
}

// The JmsClient struct provides a high level interface to manage Terminal, Account and Assets.
// It embeds the Terminal, Account and Assets struct which provide operations specific to each type.
type JmsClient struct {
	Terminal Terminal
	Account  Account
	Assets   Assets
}

// NewJmsClient is a factory function that returns a new JmsClient.
// It sets up the Terminal, Account, and Assets with the provided JmsAPIConfig,
// This makes it convenient to create a JmsClient with a common API configuration.
func NewJmsClient(api JmsAPIConfig) *JmsClient {
	return &JmsClient{
		Terminal: Terminal{
			Session: terminal.Sessions{
				API: api,
			},
		},
		Account: Account{
			Account: accouts.Account{
				API: api,
			},
		},
		Assets: Assets{
			Assets: assets.Assets{
				API: api,
			},
		},
	}
}
