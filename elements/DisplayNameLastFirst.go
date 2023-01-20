package elements

// The DisplayNameLastFirst element specifies the display name of the associated persona in the format,Last Name,First Name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaynamelastfirst
import "encoding/xml"

type DisplayNameLastFirst struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNameLastFirst) SetForMarshal() {
	D.XMLName.Local = "t:DisplayNameLastFirst"
}

func (D *DisplayNameLastFirst) GetSchema() *Schema {
	return &SchemaTypes
}
