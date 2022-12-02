package elements

// The Modify element indicates whether a client can modify a folder or item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/modify
type Modify bool

const (
	// false
	Modifyfalse Modify = false
	// true
	Modifytrue Modify = true
)
