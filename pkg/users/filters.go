package users

// UserFilter is a struct that represents the filtering options for querying account data.
// Each struct field corresponds to a possible filter, and the `url` tag specifies the URL
// query string parameter name associated with the struct field. All filters are optional,
// so empty fields will not be used in the final query.
// Allowed filters include ID, Asset, SourceID, SecretType, IP, HostName, Address,
// AssetID, Assets, Nodes, NodeID, HasSecret, Platform, Category, Type, Search, Order,
// Limit, Offset.
type UserFilter struct {
	ID             string `url:"id,omitempty"`
	Username       string `url:"username,omitempty"`
	Email          string `url:"email,omitempty"`
	Name           string `url:"name,omitempty"`
	Groups         string `url:"groups,omitempty"`
	GroupID        string `url:"group_id,omitempty"`
	ExcludeGroupID string `url:"exclude_group_id,omitempty"`
	Source         string `url:"source,omitempty"`
	OrgRoles       string `url:"org_roles,omitempty"`
	SystemRoles    string `url:"system_roles,omitempty"`
	IsActive       string `url:"is_active,omitempty"`
	Search         string `url:"search,omitempty"`
	Order          string `url:"order,omitempty"`
	Limit          int    `url:"limit,omitempty"`
	Offset         int    `url:"offset,omitempty"`
}
