package elements

// The DisplayTo element represents the display string that is used for the contents of the To box. This is the concatenated string of all To recipient display names.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displayto
import "encoding/xml"

type DisplayTo struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayTo) SetForMarshal() {
	D.XMLName.Local = "t:DisplayTo"
}

func (D *DisplayTo) GetSchema() *Schema {
	return &SchemaTypes
}
