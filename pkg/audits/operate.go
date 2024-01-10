package audits

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"gojms"
	"gojms/pkg/utils"
	"net/http"
)

// OperateLog is a structure that holds configuration for the JmsAPI.
// It contains a single field of type gojms.JmsAPIConfig which is used to make API requests.
type OperateLog struct {
	API gojms.JmsAPIConfig
}

// Get is a method on the OperateLog struct.
// It receives a string id as a parameter.
// The method checks if the id is non-empty and combines the API endpoint before making an http.Request.
// If the id is empty, the method returns an error.
// Otherwise, it creates a GET http.Request using the id to form the request URL
// and then executes the request.
// If the request execution is successful, the method returns nil. Otherwise, it returns the error from the request execution.
func (o *OperateLog) Get(id string) error {
	// check id
	if id == "" {
		return fmt.Errorf("operation log id can not empty")
	}

	// combine api endpoint
	endpoint := utils.CombineURL(o.API.Endpoints, opertateLogGetAPI)
	endpoint = fmt.Sprintf(endpoint, id)

	// make request
	req, err := o.API.MakeRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	// do request
	err = o.API.DoRequest(req, nil)
	return err

}

// List is a method on the OperateLog struct.
// It accepts a pointer to an OperateFilter object as a parameter.
// The method combines the API endpoint, makes an http.Request, sets the URL query string,
// and executes the request.
// If filter is not nil, the method generates URL string parameters from
// the filter object and appends it to the request.
// If the request is successful, the method returns nil. If not, it returns an error.
func (o *OperateLog) List(filter *OperateFilter) error {

	// combine api endpoint
	endpoint := utils.CombineURL(o.API.Endpoints, opertateLogListAPI)

	// make request
	req, err := o.API.MakeRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	// set query params
	if filter != nil {
		v, err := query.Values(filter)
		if err != nil {
			return err
		}
		req = o.API.SetQuery(req, v)
	}

	// do request
	err = o.API.DoRequest(req, nil)
	return err
}
