package terminal

import (
	"fmt"
	"github.com/MScuti/gojms/pkg/apiauth"
	"github.com/MScuti/gojms/pkg/utils"
	"github.com/google/go-querystring/query"
	"net/http"
)

// Sessions is a struct that holds configuration for the JmsAPI.
// It contains a single field of type gojms.JmsAPIConfig which is used to make API requests.
type Sessions struct {
	API apiauth.JmsAPI
}

// SessionDetailRep represents the detailed response of a session.
// It includes fields like the Id, User, Asset, UserId, AssetId, Account, AccountId,
// Protocol, Type, LoginFrom, RemoteAddr, Comment, TerminalDisplay, IsLocked,
// CommandAmount, Terminal, OrgId, OrgName, IsSuccess, IsFinished, HasReplay,
// HasCommand, CanReplay, CanJoin, CanTerminate, DateStart, and DateEnd.
// 'json' struct tags are used to map the struct fields with the json response.
type SessionDetailRep struct {
	Id        string `json:"id"`
	User      string `json:"user"`
	Asset     string `json:"asset"`
	UserId    string `json:"user_id"`
	AssetId   string `json:"asset_id"`
	Account   string `json:"account"`
	AccountId string `json:"account_id"`
	Protocol  string `json:"protocol"`
	Type      struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"type"`
	LoginFrom struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"login_from"`
	RemoteAddr      string      `json:"remote_addr"`
	Comment         interface{} `json:"comment"`
	TerminalDisplay string      `json:"terminal_display"`
	IsLocked        bool        `json:"is_locked"`
	CommandAmount   int         `json:"command_amount"`
	Terminal        struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"terminal"`
	OrgId        string `json:"org_id"`
	OrgName      string `json:"org_name"`
	IsSuccess    bool   `json:"is_success"`
	IsFinished   bool   `json:"is_finished"`
	HasReplay    bool   `json:"has_replay"`
	HasCommand   bool   `json:"has_command"`
	CanReplay    bool   `json:"can_replay"`
	CanJoin      bool   `json:"can_join"`
	CanTerminate bool   `json:"can_terminate"`
	DateStart    string `json:"date_start"`
	DateEnd      string `json:"date_end"`
}

// SessionListRep is a struct representing a paginated response for a list of sessions.
// It holds the total count of sessions,
// the URL for the next page,
// the URL for the previous page,
// and a list of detailed session representations.
type SessionListRep []SessionDetailRep

// Get is a method on the Sessions struct.
// It accepts a string id as a parameter and retrieves the session details
// associated with that id from the server.
// If id is empty, it returns immediately with an error.
// If id is non-empty, this method will make an HTTP GET request to the server
// and unmarshal the response into a SessionDetailRep object.
// If the retrieval and unmarshalling are successful, it returns the SessionDetailRep
// object along with a nil error. Otherwise, it returns nil and the associated error.
func (s *Sessions) Get(id string) (*SessionDetailRep, error) {
	// check id
	if id == "" {
		return nil, fmt.Errorf("session id can not empty")
	}

	// combine api endpoint
	endpoint := utils.CombineURL(s.API.Endpoints, sessionGetAPI)
	endpoint = fmt.Sprintf(endpoint, id)

	// make request
	req, err := s.API.MakeRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// do request
	data := &SessionDetailRep{}
	err = s.API.DoRequest(req, data)
	return data, err

}

// List is a method on the Sessions struct.
// It accepts a pointer to a SessionsFilter struct as a parameter.
// The function generates an API endpoint, makes a GET HTTP request, and sets the URL query parameters
// based on the filter argument, if it's not nil.
// Then it executes the request and captures the response data in a SessionListRep object.
// If the operations are successful, it returns a pointer to the SessionListRep object and a nil error.
// If there's an error during these operations, it returns nil and the error.
func (s *Sessions) List(filter *SessionsFilter) (*SessionListRep, error) {
	// combine api endpoint
	endpoint := utils.CombineURL(s.API.Endpoints, sessionListAPI)

	// make request
	req, err := s.API.MakeRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// set query params
	if filter != nil {
		v, err := query.Values(filter)
		if err != nil {
			return nil, err
		}
		req = s.API.SetQuery(req, v)
	}

	// do request
	data := &SessionListRep{}
	err = s.API.DoRequest(req, data)
	return data, err
}
