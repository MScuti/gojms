package gojms

import (
	"github.com/MScuti/gojms/pkg/accouts"
	"github.com/MScuti/gojms/pkg/apiauth"
	"github.com/MScuti/gojms/pkg/assets"
	"github.com/MScuti/gojms/pkg/terminal"
	"github.com/MScuti/gojms/pkg/users"
)

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

// The User struct holds the User object for user operations.
// It is used to manage and interact with users.
type User struct {
	User users.User
}

// The JmsClient struct provides a high level interface to manage Terminal, Account and Assets.
// It embeds the Terminal, Account and Assets struct which provide operations specific to each type.
type JmsClient struct {
	Terminal Terminal
	Account  Account
	Assets   Assets
	User     User
}

// JmsAKClient is a struct representing a AKClient entity in the program.
// It includes properties that a AKClient entity may have.
//
// Properties:
//
//	Terminal: This property comprises a Terminal structure for terminal session operations.
//	Account: This property holds the Account struct for account operations.
//	Assets: This property contains the Assets structure for asset management operations.
//	User: This property uses the User struct for user management operations.
//
// The struct has been developed to enable easy management and interaction with terminals, accounts,
// assets, and users.
//
// Each field in the struct is equipped to manage and interact with their respective properties.
// The struct is initialized using a factory function NewJmsAKClient which takes in a parameter of
// apiauth.JmsAKConfig type.
//
// The significance of each field is defined in their respective structs 'Terminal','Account','Assets'
// and 'User'. Each struct has corresponding methods for different operations like fetching data from
// the server, making HTTP requests, etc.
type JmsAKClient struct {
	Terminal Terminal
	Account  Account
	Assets   Assets
	User     User
}

// JmsSdkClient is a struct representing a SdkClient entity in the program.
// It includes properties that a SdkClient entity may have.
//
// Properties:
//
//	Terminal: This property comprises a Terminal structure for terminal session operations.
//	Account: This property holds the Account struct for account operations.
//	Assets: This property contains the Assets structure for asset management operations.
//	User: This property uses the User struct for user management operations.
//
// The struct has been developed to enable ease in managing and interacting with terminals, accounts,
// assets, and users.
//
// Each field in the struct is equipped to manage and interact with their respective properties.
// The struct is initialized using a factory function NewJmsSdkClient which takes in a parameter of
// apiauth.JmsSdkConfig type.
//
// The significance of each field is defined in their respective structs 'Terminal','Account','Assets'
// and 'User'. Each struct has corresponding methods for different operations like fetching data from
// the server, making HTTP requests, etc.
type JmsSdkClient struct {
	Terminal Terminal
	Account  Account
	Assets   Assets
	User     User
}

// NewJmsClient is a factory function that returns a new JmsClient.
// It sets up the Terminal, Account, and Assets with the provided JmsAPIConfig,
// This makes it convenient to create a JmsClient with a common API configuration.
func NewJmsClient(api apiauth.JmsAPIConfig) *JmsClient {
	return &JmsClient{
		Terminal: Terminal{
			Session: terminal.Sessions{
				API: &api,
			},
		},
		Account: Account{
			Account: accouts.Account{
				API: &api,
			},
		},
		Assets: Assets{
			Assets: assets.Assets{
				API: &api,
			},
		},
		User: User{
			User: users.User{API: &api},
		},
	}
}

// NewJmsAKClient is a factory function that returns a new NewJmsAKClient.
func NewJmsAKClient(api apiauth.JmsAKConfig) *JmsAKClient {
	return &JmsAKClient{
		Terminal: Terminal{
			Session: terminal.Sessions{
				API: &api,
			},
		},
		Account: Account{
			Account: accouts.Account{
				API: &api,
			},
		},
		Assets: Assets{
			Assets: assets.Assets{
				API: &api,
			},
		},
		User: User{
			User: users.User{API: &api},
		},
	}
}

// NewSdkClient is a function that initializes a new JmsSdkClient struct
// with the provided JmsSDKConfig and returns a pointer to it.
//
// Parameters:
//
//	api apiauth.JmsSDKConfig: An instance of JmsSDKConfig containing API configuration parameters.
//
// Returns:
//
//	*JmsSdkClient: A pointer to the newly initialized JmsSdkClient struct.
//
// Process:
//   - The function initializes a new JmsSdkClient struct with the provided api, setting up the Terminal, Account, Assets, and User fields.
//   - The provided JmsSDKConfig is set as the API for the Sessions field of Terminal, Account field of Account, Assets field of Assets, and User field of User.
//   - Finally, the function returns a pointer to this newly initialized JmsSdkClient struct.
func NewSdkClient(api apiauth.JmsSDKConfig) *JmsSdkClient {
	return &JmsSdkClient{
		Terminal: Terminal{
			Session: terminal.Sessions{
				API: &api,
			},
		},
		Account: Account{
			Account: accouts.Account{
				API: &api,
			},
		},
		Assets: Assets{
			Assets: assets.Assets{
				API: &api,
			},
		},
		User: User{
			User: users.User{API: &api},
		},
	}
}
