package elements

// The CreateHierarchy element indicates whether a client can create a hierarchy table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createhierarchy
type CreateHierarchy bool

const (
	// false
	CreateHierarchyfalse CreateHierarchy = false
	// true
	CreateHierarchytrue CreateHierarchy = true
)
