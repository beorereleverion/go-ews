package elements

// The ErrorDescription element describes the error that is returned in information about an item that cannot be indexed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/errordescription
import "encoding/xml"

type ErrorDescription struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *ErrorDescription) SetForMarshal() {
	E.XMLName.Local = "t:ErrorDescription"
}

func (E *ErrorDescription) GetSchema() *Schema {
	return &SchemaTypes
}
