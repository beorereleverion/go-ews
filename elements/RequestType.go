package elements

// The RequestType element identifies whether a proxy request is a cross-site or a cross-forest request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requesttype
import "encoding/xml"

type RequestType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *RequestType) SetForMarshal() {
	R.XMLName.Local = "t:RequestType"
}

func (R *RequestType) GetSchema() *Schema {
	return &SchemaTypes
}
