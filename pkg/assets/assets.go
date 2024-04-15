package assets

import (
	"fmt"
	"github.com/MScuti/gojms/pkg/apiauth"
	"github.com/MScuti/gojms/pkg/utils"
	"github.com/google/go-querystring/query"
	"net/http"
)

// The Assets struct holds the configuration for the JmsAPI.
// This structure contains a single field of type gojms.JmsAPIConfig
// which is used to make API requests.
type Assets struct {
	API apiauth.JmsAPI
}

// AssetDetailRep represents the details of an asset.
// It includes fields such as Id, Name, Address, Comment, Domain, Platform, Nodes, Labels, Protocols, NodesDisplay, Accounts,
// Category, Type, Connectivity, AutoConfig, CreatedBy, OrgId, OrgName, GatheredInfo, SpecInfo, IsActive, DateVerified, and DateCreated.
// The 'json' struct tags are used to map the struct fields with the json response.
type AssetDetailRep struct {
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	Address  string      `json:"address"`
	Comment  string      `json:"comment"`
	Domain   interface{} `json:"domain"`
	Platform struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"platform"`
	Nodes []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"nodes"`
	Labels    []interface{} `json:"labels"`
	Protocols []struct {
		Name string `json:"name"`
		Port int    `json:"port"`
	} `json:"protocols"`
	NodesDisplay []string `json:"nodes_display"`
	Accounts     []struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Username   string `json:"username"`
		SecretType struct {
			Value string `json:"value"`
			Label string `json:"label"`
		} `json:"secret_type"`
		CreatedBy string      `json:"created_by"`
		Comment   string      `json:"comment"`
		SuFrom    interface{} `json:"su_from"`
		Version   int         `json:"version"`
		Source    struct {
			Value string `json:"value"`
			Label string `json:"label"`
		} `json:"source"`
		SourceId     string `json:"source_id"`
		Connectivity struct {
			Value string `json:"value"`
			Label string `json:"label"`
		} `json:"connectivity"`
		Privileged  bool   `json:"privileged"`
		IsActive    bool   `json:"is_active"`
		DateCreated string `json:"date_created"`
		DateUpdated string `json:"date_updated"`
	} `json:"accounts"`
	Category struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"category"`
	Type struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"type"`
	Connectivity struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"connectivity"`
	AutoConfig struct {
		SuEnabled      bool `json:"su_enabled"`
		DomainEnabled  bool `json:"domain_enabled"`
		AnsibleEnabled bool `json:"ansible_enabled"`
		Id             int  `json:"id"`
		AnsibleConfig  struct {
			AnsibleConnection string `json:"ansible_connection"`
		} `json:"ansible_config"`
		PingEnabled bool   `json:"ping_enabled"`
		PingMethod  string `json:"ping_method"`
		PingParams  struct {
		} `json:"ping_params"`
		GatherFactsEnabled bool   `json:"gather_facts_enabled"`
		GatherFactsMethod  string `json:"gather_facts_method"`
		GatherFactsParams  struct {
		} `json:"gather_facts_params"`
		ChangeSecretEnabled bool   `json:"change_secret_enabled"`
		ChangeSecretMethod  string `json:"change_secret_method"`
		ChangeSecretParams  struct {
		} `json:"change_secret_params"`
		PushAccountEnabled bool   `json:"push_account_enabled"`
		PushAccountMethod  string `json:"push_account_method"`
		PushAccountParams  struct {
			Home   string `json:"home"`
			Sudo   string `json:"sudo"`
			Shell  string `json:"shell"`
			Groups string `json:"groups"`
		} `json:"push_account_params"`
		VerifyAccountEnabled bool   `json:"verify_account_enabled"`
		VerifyAccountMethod  string `json:"verify_account_method"`
		VerifyAccountParams  struct {
		} `json:"verify_account_params"`
		GatherAccountsEnabled bool   `json:"gather_accounts_enabled"`
		GatherAccountsMethod  string `json:"gather_accounts_method"`
		GatherAccountsParams  struct {
		} `json:"gather_accounts_params"`
		Platform int `json:"platform"`
	} `json:"auto_config"`
	CreatedBy    string `json:"created_by"`
	OrgId        string `json:"org_id"`
	OrgName      string `json:"org_name"`
	GatheredInfo struct {
	} `json:"gathered_info"`
	SpecInfo struct {
	} `json:"spec_info"`
	IsActive     bool        `json:"is_active"`
	DateVerified interface{} `json:"date_verified"`
	DateCreated  string      `json:"date_created"`
}

// AssetListRep is a type that represents a list of AssetDetailRep objects.
// It's basically a slice of AssetDetailRep struct instances which represent multiple asset details.
type AssetListRep []AssetDetailRep

// Get is a method on the Assets struct.
// It takes a string id as a parameter and retrieves the asset details
// associated with that id from the server.
// If the id is empty, it returns immediately with an error.
// If the id is non-empty, this method will make an HTTP GET request to the server
// and unmarshal the response into an AssetDetailRep object.
// If the retrieval and unmarshalling are successful, it returns the AssetDetailRep
// object along with a nil error. Otherwise, it returns nil and the associated error.
func (s *Assets) Get(id string) (*AssetDetailRep, error) {
	// check id
	if id == "" {
		return nil, fmt.Errorf("session id can not empty")
	}

	// combine api endpoint
	endpoint := utils.CombineURL(s.API.GetEndpoint(), assetsGetAPI)
	endpoint = fmt.Sprintf(endpoint, id)

	// make request
	req, err := s.API.MakeRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// do request
	data := &AssetDetailRep{}
	err = s.API.DoRequest(req, data)
	return data, err

}

// List is a method on the Assets struct.
// It accepts a pointer to an AssetFilter object as a parameter.
// The function generates an API endpoint, makes a GET HTTP request, and sets the URL query parameters
// based on the filter argument, if it's not null.
// Then it executes the request and captures the response data in an AssetListRep object.
// If the operations are successful, it returns a pointer to the AssetListRep object and a nil error.
// If there's an error during these operations, it returns nil and the relevant error.
func (s *Assets) List(filter *AssetFilter) (*AssetListRep, error) {
	// combine api endpoint
	endpoint := utils.CombineURL(s.API.GetEndpoint(), assetsListAPI)

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
	data := &AssetListRep{}
	err = s.API.DoRequest(req, data)
	return data, err
}
