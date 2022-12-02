package elements

// The DelegateUser element identifies a single delegate to add or update in a mailbox or a delegate returned in a delegate management response. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delegateuser
type DelegateUser struct {
	// The DelegatePermissions element contains the delegate permission-level settings for a user. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DelegatePermissions *DelegatePermissions `xml:"t:DelegatePermissions"`
	// The ReceiveCopiesOfMeetingMessages element indicates whether a delegate receives copies of meeting-related messages that are addressed to the principal. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	ReceiveCopiesOfMeetingMessages *ReceiveCopiesOfMeetingMessages `xml:"t:ReceiveCopiesOfMeetingMessages"`
	// The UserId element identifies a delegate user or a user who has folder access permissions.
	UserId *UserId `xml:"t:UserId"`
	// The ViewPrivateItems element indicates whether a delegate user or client application has permission to view private items in the principal's mailbox.
	ViewPrivateItems *ViewPrivateItems `xml:"t:ViewPrivateItems"`
}
