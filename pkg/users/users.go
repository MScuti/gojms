package users

import (
	"fmt"
	"github.com/MScuti/gojms/pkg/apiauth"
	"github.com/MScuti/gojms/pkg/utils"
	"github.com/google/go-querystring/query"
	"net/http"
)

// The User struct holds the configuration for the JmsAPI.
// This structure contains a single field of type gojms.JmsAPIConfig
// which is used to make API requests.
type User struct {
	API apiauth.JmsAPIConfig
}

// UserDetailRep serves as a representation for account details.
// It features an extensive set of fields correlating to a broad range of account information,
// such as user id, username, authentication, roles, and more.
// Fields are tagged with 'json' specifying the JSON payload equivalent name.
// Empty field means the account detail is not available / not applicable.
type UserDetailRep struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Wechat   string `json:"wechat"`
	Phone    string `json:"phone"`
	MfaLevel struct {
		Value int    `json:"value"`
		Label string `json:"label"`
	} `json:"mfa_level"`
	Source struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"source"`
	WecomId     interface{}   `json:"wecom_id"`
	DingtalkId  interface{}   `json:"dingtalk_id"`
	FeishuId    interface{}   `json:"feishu_id"`
	CreatedBy   string        `json:"created_by"`
	UpdatedBy   string        `json:"updated_by"`
	Comment     string        `json:"comment"`
	IsSuperuser bool          `json:"is_superuser"`
	IsOrgAdmin  bool          `json:"is_org_admin"`
	AvatarUrl   string        `json:"avatar_url"`
	Groups      []interface{} `json:"groups"`
	SystemRoles []struct {
		Id          string `json:"id"`
		DisplayName string `json:"display_name"`
	} `json:"system_roles"`
	OrgRoles []struct {
		Id          string `json:"id"`
		DisplayName string `json:"display_name"`
		Name        string `json:"name"`
	} `json:"org_roles"`
	PasswordStrategy struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"password_strategy"`
	IsServiceAccount        bool   `json:"is_service_account"`
	IsValid                 bool   `json:"is_valid"`
	IsExpired               bool   `json:"is_expired"`
	IsActive                bool   `json:"is_active"`
	IsOtpSecretKeyBound     bool   `json:"is_otp_secret_key_bound"`
	CanPublicKeyAuth        bool   `json:"can_public_key_auth"`
	MfaEnabled              bool   `json:"mfa_enabled"`
	NeedUpdatePassword      bool   `json:"need_update_password"`
	MfaForceEnabled         bool   `json:"mfa_force_enabled"`
	IsFirstLogin            bool   `json:"is_first_login"`
	LoginBlocked            bool   `json:"login_blocked"`
	DateExpired             string `json:"date_expired"`
	DateJoined              string `json:"date_joined"`
	LastLogin               string `json:"last_login"`
	DateUpdated             string `json:"date_updated"`
	DatePasswordLastUpdated string `json:"date_password_last_updated"`
}

// UserListRep is a type that represents a list of AccountDetailRep objects.
// It's essentially a slice of AccountDetailRep struct instances representing multiple accounts details.
type UserListRep []UserDetailRep

// Get is a method on the Account struct.
// It accepts a string id as a parameter and retrieves the account details
// associated with that id from the server.
// If id is empty, it returns immediately with an error.
// This method will make an HTTP GET request to the server and
// unmarshal the response into an AccountDetailRep object.
// If the retrieval and unmarshalling are successful, it returns the AccountDetailRep
// object along with a nil error. Otherwise, it returns nil and the associated error.
func (u *User) Get(id string) (*UserDetailRep, error) {
	// check id
	if id == "" {
		return nil, fmt.Errorf("session id can not empty")
	}

	// combine api endpoint
	endpoint := utils.CombineURL(u.API.Endpoints, accountsGetAPI)
	endpoint = fmt.Sprintf(endpoint, id)

	// make request
	req, err := u.API.MakeRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// do request
	data := &UserDetailRep{}
	err = u.API.DoRequest(req, data)
	return data, err
}

// List is a method on the Account struct.
// It accepts a pointer to an AccountFilter object as a parameter.
// The function generates an API endpoint, makes a GET HTTP request, and sets the URL query parameters
// based on the filter argument, if it's not null.
// Then it executes the request and captures the response data in an AccountListRep object.
// If the operations are successful, it returns a pointer to the AccountListRep object and a nil error.
// If there's an error during these operations, it returns nil and the error.
func (u *User) List(filter *UserFilter) (*UserListRep, error) {
	// combine api endpoint
	endpoint := utils.CombineURL(u.API.Endpoints, accountsListAPI)

	// make request
	req, err := u.API.MakeRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// set query params
	if filter != nil {
		v, err := query.Values(filter)
		if err != nil {
			return nil, err
		}
		req = u.API.SetQuery(req, v)
	}

	// do request
	data := &UserListRep{}
	err = u.API.DoRequest(req, data)
	return data, err
}
