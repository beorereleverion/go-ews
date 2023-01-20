package elements

// The OriginalDisplayName element contains the original display name associated with an email address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/originaldisplayname
import "encoding/xml"

type OriginalDisplayName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (O *OriginalDisplayName) SetForMarshal() {
	O.XMLName.Local = "t:OriginalDisplayName"
}

func (O *OriginalDisplayName) GetSchema() *Schema {
	return &SchemaTypes
}
