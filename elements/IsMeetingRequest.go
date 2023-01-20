package elements

// The IsMeetngRequest element indicates whether incoming messages must be a meeting request in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismeetingrequest
import "encoding/xml"

type IsMeetingRequest struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsMeetingRequestfalse bool = false
	// true
	IsMeetingRequesttrue bool = true
)

func (I *IsMeetingRequest) SetForMarshal() {
	I.XMLName.Local = "m:IsMeetingRequest"
}

func (I *IsMeetingRequest) GetSchema() *Schema {
	return &SchemaMessages
}
