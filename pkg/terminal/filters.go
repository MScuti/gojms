package terminal

// SessionsFilter represents the filters that can be applied when querying session data.
// The `url` struct tag is used to specify the URL query string parameter name associated with the struct field.
// User: Filter sessions by user.
// Asset: Filter sessions by asset.
// Account: Filter sessions by account.
// RemoteAddr: Filter sessions by remote address.
// Protocol: Filter sessions by protocol.
// IsFinished: Filter sessions by finish status.
// LoginFrom: Filter sessions by login source.
// Terminal: Filter sessions by terminal.
// Search: General search across multiple fields.
// Order: Specify the order of the returned data.
// Limit: Limit the number of returned items.
// Offset: Offset the start point of the returned data. Useful for paginated data retrieval.
type SessionsFilter struct {
	User       string `url:"user"`
	Asset      string `url:"asset"`
	Account    string `url:"account"`
	RemoteAddr string `url:"remote_addr"`
	Protocol   string `url:"protocol"`
	IsFinished string `url:"is_finished"`
	LoginFrom  string `url:"login_from"`
	Terminal   string `url:"terminal"`
	Search     string `url:"search"`
	Order      string `url:"order"`
	Limit      int    `url:"limit"`
	Offset     int    `url:"offset"`
	DateFrom   string `url:"date_from"`
	DateTo     string `url:"date_to"`
}
