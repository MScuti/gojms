package users

// UserFilter is a struct that represents the filtering options for querying account data.
// Each struct field corresponds to a possible filter, and the `url` tag specifies the URL
// query string parameter name associated with the struct field. All filters are optional,
// so empty fields will not be used in the final query.
// Allowed filters include ID, Asset, SourceID, SecretType, IP, HostName, Address,
// AssetID, Assets, Nodes, NodeID, HasSecret, Platform, Category, Type, Search, Order,
// Limit, Offset.
type UserFilter struct {
	ID             string `url:"id"`
	Username       string `url:"username"`
	Email          string `url:"email"`
	Name           string `url:"name"`
	Groups         string `url:"groups"`
	GroupID        string `url:"group_id"`
	ExcludeGroupID string `url:"exclude_group_id"`
	Source         string `url:"source"`
	OrgRoles       string `url:"org_roles"`
	SystemRoles    string `url:"system_roles"`
	IsActive       string `url:"is_active"`
	Search         string `url:"search"`
	Order          string `url:"order"`
	Limit          int    `url:"limit"`
	Offset         int    `url:"offset"`
}
