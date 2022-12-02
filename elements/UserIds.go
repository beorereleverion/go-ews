package elements

// The UserIds element contains an array of delegate users to get or remove from a principal's mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userids
type UserIds struct {
	// The UserId element identifies a delegate user or a user who has folder access permissions.
	UserId *UserId `xml:"t:UserId"`
}
