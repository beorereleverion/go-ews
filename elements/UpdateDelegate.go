package elements

// The UpdateDelegate element defines a request to update delegates in a mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatedelegate
import "encoding/xml"

type UpdateDelegate struct {
	XMLName xml.Name

	// The DelegateUser element identifies a single delegate to add or update in a mailbox or a delegate returned in a delegate management response. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DelegateUser *DelegateUser `xml:"DelegateUser"`
	// The DelegateUsers element contains the identities of delegates to add to or update in a mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DelegateUsers *DelegateUsers `xml:"DelegateUsers"`
	// The DeliverMeetingRequests element defines how meeting requests are handled between the delegate and the principal. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DeliverMeetingRequests *DeliverMeetingRequests `xml:"DeliverMeetingRequests"`
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (U *UpdateDelegate) SetForMarshal() {
	U.XMLName.Local = "m:UpdateDelegate"
}

func (U *UpdateDelegate) GetSchema() *Schema {
	return &SchemaMessages
}
