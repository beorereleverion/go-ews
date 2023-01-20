package elements

// The DisplayNamePrefix element specifies the prefix of the display name of the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaynameprefix
import "encoding/xml"

type DisplayNamePrefix struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNamePrefix) SetForMarshal() {
	D.XMLName.Local = "t:DisplayNamePrefix"
}

func (D *DisplayNamePrefix) GetSchema() *Schema {
	return &SchemaTypes
}
