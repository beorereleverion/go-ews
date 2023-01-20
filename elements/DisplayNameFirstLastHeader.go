package elements

// The DisplayNameFirstLastHeader element specifies the header for the display name, first name first.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaynamefirstlastheader
import "encoding/xml"

type DisplayNameFirstLastHeader struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNameFirstLastHeader) SetForMarshal() {
	D.XMLName.Local = "t:DisplayNameFirstLastHeader"
}

func (D *DisplayNameFirstLastHeader) GetSchema() *Schema {
	return &SchemaTypes
}
