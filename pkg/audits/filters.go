package audits

// OperateFilter represents a filter for operation requests.
// filter for api: /audits/operate-logs/
type OperateFilter struct {
	User         string `url:"user"`
	Action       string `url:"action"`
	ResourceType string `url:"resource_type"`
	Resource     string `url:"resource"`
	RemoteAddr   string `url:"remote_addr"`
	Search       string `url:"search"`
	Order        string `url:"order"`
	Limit        int    `url:"limit"`
	Offset       int    `url:"offset"`
}
