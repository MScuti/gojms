package assets

type AssetFilter struct {
	ID                    string `url:"id"`
	Name                  string `url:"name"`
	Address               string `url:"address"`
	IsActive              string `url:"is_active"`
	Labels                string `url:"labels"`
	Type                  string `url:"type"`
	Category              string `url:"category"`
	Platform              string `url:"platform"`
	Domain                string `url:"domain"`
	Protocols             string `url:"protocols"`
	DomainEnabled         string `url:"domain_enabled"`
	PingEnabled           string `url:"ping_enabled"`
	GatherFactsEnabled    string `url:"gather_facts_enabled"`
	ChangeSecretEnabled   string `url:"change_secret_enabled"`
	PushAccountEnabled    string `url:"push_account_enabled"`
	VerifyAccountEnabled  string `url:"verify_account_enabled"`
	GatherAccountsEnabled string `url:"gather_accounts_enabled"`
	Search                string `url:"search"`
	Order                 string `url:"order"`
	Limit                 int    `url:"limit"`
	Offset                int    `url:"offset"`
}
