package elements

// The ResponseType element represents the type of recipient response that is received for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsetype
import "encoding/xml"

type ResponseType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *ResponseType) SetForMarshal() {
	R.XMLName.Local = "t:ResponseType"
}

func (R *ResponseType) GetSchema() *Schema {
	return &SchemaTypes
}
