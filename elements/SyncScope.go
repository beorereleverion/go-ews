package elements

// The SyncScope element specifies whether just items or items and folder associated information are returned in a synchronization response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncscope
type SyncScope string

const (
	// NormalAndAssociatedItems
	SyncScopeNormalAndAssociatedItems SyncScope = `NormalAndAssociatedItems`
	// NormalItems
	SyncScopeNormalItems SyncScope = `NormalItems`
)
