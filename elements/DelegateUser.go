package elements

// The DelegateUser element identifies a single delegate to add or update in a mailbox or a delegate returned in a delegate management response. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delegateuser
import "encoding/xml"

type DelegateUser struct {
	XMLName xml.Name

	// The DelegatePermissions element contains the delegate permission-level settings for a user. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DelegatePermissions *DelegatePermissions `xml:"DelegatePermissions"`
	// The ReceiveCopiesOfMeetingMessages element indicates whether a delegate receives copies of meeting-related messages that are addressed to the principal. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	ReceiveCopiesOfMeetingMessages *ReceiveCopiesOfMeetingMessages `xml:"ReceiveCopiesOfMeetingMessages"`
	// The UserId element identifies a delegate user or a user who has folder access permissions.
	UserId *UserId `xml:"UserId"`
	// The ViewPrivateItems element indicates whether a delegate user or client application has permission to view private items in the principal's mailbox.
	ViewPrivateItems *ViewPrivateItems `xml:"ViewPrivateItems"`
}

func (D *DelegateUser) SetForMarshal() {
	D.XMLName.Local = "t:DelegateUser"
}

func (D *DelegateUser) GetSchema() *Schema {
	return &SchemaTypes
}
