package elements

// The BusinessName element specifies the name of a business.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/businessname
import "encoding/xml"

type BusinessName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (B *BusinessName) SetForMarshal() {
	B.XMLName.Local = "t:BusinessName"
}

func (B *BusinessName) GetSchema() *Schema {
	return &SchemaTypes
}
