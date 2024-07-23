package apiauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cyberark/conjur-api-go/conjurapi"
	"gopkg.in/twindagger/httpsig.v1"
	"io"
	"net/http"
	"net/url"
	"os"
)

// JmsSDKConfig is a struct that handles the configuration for JmsSDK.
//
// Properties:
//
//	Endpoints: A string field that represents the API endpoints URLs being used by the JmsSDK.
//	Debug: A boolean field indicating if debug mode is enabled. When true, response bodies are printed to the console.
//	ConjurFileName: The name of the Conjur file, typically used for API authorization.
//
// The struct fields are serializable to JSON with respective tags provided.
//
// Major methods include:
//
//	SignReq: Reads environment variables, initializes the Conjur client, retrieves ak and sk,
//	and uses them to sign an HTTP request.
//
//	MakeRequest: Processes the request body data, makes a new HTTP request, sets the header,
//	and signs the request using the SignReq method.
//
//	DoRequest: Unmarshals the response body of an HTTP request into the result parameter.
//
//	SetQuery: Sets the provided query parameters (given as url.Values) on an HTTP request.
//
//	GetEndpoint: Getter method for the `Endpoints` field of the JmsSDKConfig struct.
//
// Each field in the struct corresponds to a configuration setting for the JmsSDK.
type JmsSDKConfig struct {
	Endpoints      string `json:"endpoints"`
	Debug          bool   `json:"debug"`
	ConjurFileName string `json:"conjur_file_name"`
}

// SignReq is a method that signs an HTTP request. It reads required environment variables,
// initializes a Conjura client, retrieves AccessKey and SecretKey, and uses these to sign the request.
//
// Parameters:
//
//	r *http.Request: The HTTP request that needs to be signed.
//
// Returns:
//
//	error: Returns an error if any environmental variable is not found or if there's any problem
//	in initializing the Conjura client, retrieving the AccessKey or SecretKey, or signing the request.
//	Otherwise, it returns nil.
//
// Process:
//   - The method first reads the environment variables 'ConjurEnvName', 'AKPath' and 'SKPath'.
//     If any of these are not found, it returns an error.
//   - If the 'ConjurEnvName' exists, it reads the Conjura file from the environment.
//     If the file does not exist, it returns an error.
//   - The method then initializes a Conjura client. If an error occurs during this process, it returns an error.
//   - It retrieves the AccessKey and SecretKey using the initialized client.
//     If an error occurs in retrieving these, it returns an error.
//   - Finally, it signs the request using the retrieved AccessKey and SecretKey.
//     If any error occurs during this process, it returns the error.
func (j *JmsSDKConfig) SignReq(r *http.Request) error {

	// read envs
	conjurFile := os.Getenv("ConjurEnvName")
	if conjurFile == "" {
		return errors.New("ConjurEnvName not found")
	}
	akPath := os.Getenv("AKPath")
	if akPath == "" {
		return errors.New("AKPath not found")
	}
	skPath := os.Getenv("SKPath")
	if skPath == "" {
		return errors.New("SKPath not found")
	}
	cFile := os.Getenv(conjurFile)
	if cFile == "" {
		return errors.New("CONJUR_AUTHN_TOKEN_FILE not found")
	}

	// init conjur client
	config, err := conjurapi.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading conjur config: %v", err)
	}
	cClient, err := conjurapi.NewClientFromTokenFile(config, cFile)
	if err != nil {
		return fmt.Errorf("error creating conjur client: %v", err)
	}

	// retrieve ak and sk
	ak, err := cClient.RetrieveSecret(akPath)
	if err != nil {
		return fmt.Errorf("error retrieving ak: %v", err)
	}
	sk, err := cClient.RetrieveSecret(skPath)
	if err != nil {
		return fmt.Errorf("error retrieving sk: %v", err)
	}

	// sign request
	signer, err := httpsig.NewRequestSigner(string(ak), string(sk), "hmac-sha256")
	if err != nil {
		return err
	}
	return signer.SignRequest(r, []string{"(request-target)", "date"}, nil)
}

// MakeRequest is a method that prepares and returns an HTTP request for given parameters.
// It marshals the body data and applies the result as request body. It also signs the request.
//
// Parameters:
//
//	method string: The HTTP method (GET, POST, PUT, DELETE).
//
//	endpoint string: The API endpoint URL.
//
//	body interface{}: The request body data. It can be any data type that can be marshalled into JSON.
//
// Returns:
//
//	*http.Request: The prepared HTTP request.
//
//	error: Returns an error if there is an issue with marshalling the provided body data,
//	creating the new HTTP request or signing the request. Otherwise, it returns nil.
//
// Process:
//   - The method first checks if the provided body data is not nil.
//     If it's not nil, it tries to marshal it into JSON format. If an error occurs during this process, it returns the error.
//   - It then creates a new HTTP request with the provided method and endpoint, and the marshalled body data.
//     If an error occurs during this process, it returns the error.
//   - It sets the 'Content-Type' of the request header to 'application/json'.
//   - It calls the SignReq method to sign the request. If an error occurs during this process, it returns the error.
//   - Finally, if everything is successful, it returns the prepared HTTP request.
func (j *JmsSDKConfig) MakeRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	// process body data
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("make request error marshal body error: %s", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	// make request
	req, err := http.NewRequest(method, endpoint, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("make new request error: %s", err)
	}

	// set header
	req.Header.Set("Content-Type", "application/json")
	err = j.SignReq(req)
	if err != nil {
		return nil, fmt.Errorf("sign request error: %s", err)
	}
	return req, nil
}

// DoRequest is a method that sends the provided HTTP request and unmarshals the response body into the provided result parameter.
// If the Debug field of the JmsSDKConfig struct is true, it also prints the response body to the console.
//
// Parameters:
//
//	req *http.Request: The HTTP request to be sent.
//
//	result interface{}: The structure in which the response body should be unmarshaled.
//	If this parameter is nil, it does not attempt to unmarshal the response body.
//
// Returns:
//
//	error: Returns an error if there's an issue sending the request, reading the response body,
//	closing the response body, the status code of the response is not in the 200-399 range,
//	or there's an issue unmarshaling the response body. Otherwise, it returns nil.
//
// Process:
//   - The method first sends the provided HTTP request. If an error occurs during this process, it returns the error.
//   - It reads the response body. If an error occurs during this process, it returns the error.
//   - The method ensures that the response body is closed when all the processing on it has been done.
//   - If the Debug field of the JmsSDKConfig struct is true, it prints the response body to the console.
//   - It checks the status code of the response. If it is not in the 200-399 range, it
func (j *JmsSDKConfig) DoRequest(req *http.Request, result interface{}) error {
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	// set debug
	if j.Debug {
		fmt.Printf("response body: %s\n", body)
	}

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

// SetQuery is a method that sets the provided query parameters on the given HTTP request.
// The provided query parameters must be of url.Values type.
// It then returns the modified HTTP request with the set query parameters.
//
// Parameters:
//
//	req *http.Request: The HTTP request on which to set the query parameters.
//
//	v url.Values: The query parameters to be set on the HTTP request. They must be of url.Values type.
//
// Returns:
//
//	*http.Request: Returns the modified HTTP request with the set query parameters.
//
// Process:
//   - The method first encodes the provided query parameters.
//   - It then sets the encoded query parameters as the RawQuery of the URL of the provided HTTP request.
//   - Finally, it returns the modified HTTP request.
func (j *JmsSDKConfig) SetQuery(req *http.Request, v url.Values) *http.Request {
	// set query
	req.URL.RawQuery = v.Encode()
	return req
}

// GetEndpoint is a method that returns the Endpoints field of the JmsSDKConfig struct.
//
// Parameters:
//
//   - None
//
// Returns:
//
//	string: The API endpoint URLs stored in the Endpoints field of the JmsSDKConfig struct.
//
// Process:
//   - The method simply returns the Endpoints field of the JmsSDKConfig struct.
func (j *JmsSDKConfig) GetEndpoint() string {
	return j.Endpoints
}
