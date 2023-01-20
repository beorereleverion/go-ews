package elements

// The DeliverMeetingRequests element defines how meeting requests are handled between the delegate and the principal. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delivermeetingrequests
import "encoding/xml"

type DeliverMeetingRequests struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// DelegatesAndMe
	DeliverMeetingRequestsDelegatesAndMe string = `DelegatesAndMe`
	// DelegatesAndSendInformationToMe
	DeliverMeetingRequestsDelegatesAndSendInformationToMe string = `DelegatesAndSendInformationToMe`
	// DelegatesOnly
	DeliverMeetingRequestsDelegatesOnly string = `DelegatesOnly`
	// NoForward
	DeliverMeetingRequestsNoForward string = `NoForward`
)

func (D *DeliverMeetingRequests) SetForMarshal() {
	D.XMLName.Local = "m:DeliverMeetingRequests"
}

func (D *DeliverMeetingRequests) GetSchema() *Schema {
	return &SchemaMessages
}
