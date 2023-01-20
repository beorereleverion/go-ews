package elements

// The IsMeetngResponsequest element indicates whether incoming messages must be a meeting response in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismeetingresponse
import "encoding/xml"

type IsMeetingResponse struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsMeetingResponsefalse bool = false
	// true
	IsMeetingResponsetrue bool = true
)

func (I *IsMeetingResponse) SetForMarshal() {
	I.XMLName.Local = "m:IsMeetingResponse"
}

func (I *IsMeetingResponse) GetSchema() *Schema {
	return &SchemaMessages
}
