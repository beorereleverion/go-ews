package elements

// The DelegateUsers element contains the identities of delegates to add to or update in a mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delegateusers
type DelegateUsers struct {
	// The DelegateUser element identifies a single delegate to add or update in a mailbox or a delegate returned in a delegate management response. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DelegateUser *DelegateUser `xml:"t:DelegateUser"`
}
