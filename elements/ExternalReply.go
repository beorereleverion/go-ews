package elements

// The ExternalReply element contains the out of office (OOF) response that is sent to addresses outside the recipient's domain or trusted domains.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externalreply
import "encoding/xml"

type ExternalReply struct {
	XMLName xml.Name

	// The Message element contains the out of Office (OOF) response.
	Message *MessageAvailability `xml:"Message"`
}

func (E *ExternalReply) SetForMarshal() {
	E.XMLName.Local = "t:ExternalReply"
}

func (E *ExternalReply) GetSchema() *Schema {
	return &SchemaTypes
}
