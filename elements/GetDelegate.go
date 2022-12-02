package elements

// The GetDelegate element defines a request to get information about delegates to a mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getdelegate
type GetDelegate struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
	// The UserIds element contains an array of delegate users to get or remove from a principal's mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	UserIds *UserIds `xml:"m:UserIds"`
	// Indicates whether the response contains permission settings for each delegate user.
	IncludePermissions *string `xml:"IncludePermissions,attr"`
}

const (
	// Delegate user permissions are returned in addition to the delegate user information that is returned in the UserId element.
	GetDelegateTrue = `True`
	// UserId information is returned.
	GetDelegateFalse = `False`
)
