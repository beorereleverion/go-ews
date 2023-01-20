package elements

// The MyResponseType element contains the status of or response to a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/myresponsetype
import "encoding/xml"

type MyResponseType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MyResponseType) SetForMarshal() {
	M.XMLName.Local = "t:MyResponseType"
}

func (M *MyResponseType) GetSchema() *Schema {
	return &SchemaTypes
}
