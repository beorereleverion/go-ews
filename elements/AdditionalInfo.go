package elements

// The AdditionalInfo element specifies additional information about the hold status of a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/additionalinfo
import "encoding/xml"

type AdditionalInfo struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *AdditionalInfo) SetForMarshal() {
	A.XMLName.Local = "t:AdditionalInfo"
}

func (A *AdditionalInfo) GetSchema() *Schema {
	return &SchemaTypes
}
