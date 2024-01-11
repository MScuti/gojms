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

// NewJmsClient is a factory function that returns a new JmsClient.
// It sets up the Terminal, Account, and Assets with the provided JmsAPIConfig,
// This makes it convenient to create a JmsClient with a common API configuration.
func NewJmsClient(api apiauth.JmsAPIConfig) *JmsClient {
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
		User: User{
			User: users.User{API: api},
		},
	}
}
