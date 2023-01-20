package elements

// The HoldId element contains the mailbox hold identifier.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/holdid
import "encoding/xml"

type HoldId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (H *HoldId) SetForMarshal() {
	H.XMLName.Local = "t:HoldId"
}

func (H *HoldId) GetSchema() *Schema {
	return &SchemaTypes
}
