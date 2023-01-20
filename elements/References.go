package elements

// The References element represents the Usenet header that is used to associate replies with the original messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/references
import "encoding/xml"

type References struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *References) SetForMarshal() {
	R.XMLName.Local = "t:References"
}

func (R *References) GetSchema() *Schema {
	return &SchemaTypes
}
