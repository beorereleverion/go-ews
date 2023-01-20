package elements

// The DisplayCc element represents the display string that is used for the contents of the Cc box. This is the concatenated string of all Cc recipient display names.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaycc
import "encoding/xml"

type DisplayCc struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayCc) SetForMarshal() {
	D.XMLName.Local = "t:DisplayCc"
}

func (D *DisplayCc) GetSchema() *Schema {
	return &SchemaTypes
}
