package elements

// The InternalReply element contains the out of office (OOF) response sent to other users in the user's domain or trusted domains.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/internalreply
import "encoding/xml"

type InternalReply struct {
	XMLName xml.Name

	// The Message element contains the out of Office (OOF) response.
	Message *MessageAvailability `xml:"Message"`
}

func (I *InternalReply) SetForMarshal() {
	I.XMLName.Local = "t:InternalReply"
}

func (I *InternalReply) GetSchema() *Schema {
	return &SchemaTypes
}
