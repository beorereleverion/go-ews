package elements

// The ContentOwner element specifies the name of the content owner.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contentowner
import "encoding/xml"

type ContentOwner struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ContentOwner) SetForMarshal() {
	C.XMLName.Local = "t:ContentOwner"
}

func (C *ContentOwner) GetSchema() *Schema {
	return &SchemaTypes
}
