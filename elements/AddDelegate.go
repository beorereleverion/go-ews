package elements

// The AddDelegate element defines a request to add delegates to a mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/adddelegate
import "encoding/xml"

type AddDelegate struct {
	XMLName xml.Name

	// The DelegateUsers element contains the identities of delegates to add to or update in a mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DelegateUsers *DelegateUsers `xml:"DelegateUsers"`
	// The DeliverMeetingRequests element defines how meeting requests are handled between the delegate and the principal. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DeliverMeetingRequests *DeliverMeetingRequests `xml:"DeliverMeetingRequests"`
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (A *AddDelegate) SetForMarshal() {
	A.XMLName.Local = "m:AddDelegate"
}

func (A *AddDelegate) GetSchema() *Schema {
	return &SchemaMessages
}
