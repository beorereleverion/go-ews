package elements

// The DisplayNameLastFirstHeader element specifies the header for the display name, last name first.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaynamelastfirstheader
import "encoding/xml"

type DisplayNameLastFirstHeader struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNameLastFirstHeader) SetForMarshal() {
	D.XMLName.Local = "t:DisplayNameLastFirstHeader"
}

func (D *DisplayNameLastFirstHeader) GetSchema() *Schema {
	return &SchemaTypes
}
