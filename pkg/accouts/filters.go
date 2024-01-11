package accouts

// AccountFilter is a struct that represents the filtering options for querying account data.
// Each struct field corresponds to a possible filter, and the `url` tag specifies the URL
// query string parameter name associated with the struct field. All filters are optional,
// so empty fields will not be used in the final query.
// Allowed filters include ID, Asset, SourceID, SecretType, IP, HostName, Address,
// AssetID, Assets, Nodes, NodeID, HasSecret, Platform, Category, Type, Search, Order,
// Limit, Offset.
type AccountFilter struct {
	ID         string `url:"id"`
	Asset      string `url:"asset"`
	SourceID   string `url:"source_id"`
	SecretType string `url:"secret_type"`
	IP         string `url:"ip"`
	HostName   string `url:"hostname"`
	Address    string `url:"address"`
	AssetID    string `url:"asset_id"`
	Assets     string `url:"assets"`
	Nodes      string `url:"nodes"`
	NodeID     string `url:"node_id"`
	HasSecret  string `url:"has_secret"`
	Platform   string `url:"platform"`
	Category   string `url:"category"`
	Type       string `url:"type"`
	Search     string `url:"search"`
	Order      string `url:"order"`
	Limit      int    `url:"limit"`
	Offset     int    `url:"offset"`
}
