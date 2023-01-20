package elements

// The DisplayNameFirstLast element specifies the display name of the associated persona in the format,First Name,Last Name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaynamefirstlast
import "encoding/xml"

type DisplayNameFirstLast struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNameFirstLast) SetForMarshal() {
	D.XMLName.Local = "t:DisplayNameFirstLast"
}

func (D *DisplayNameFirstLast) GetSchema() *Schema {
	return &SchemaTypes
}
