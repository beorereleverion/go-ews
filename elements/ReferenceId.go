package elements

// The ReferenceId element specifies the reference identifier for the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/referenceid
import "encoding/xml"

type ReferenceId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *ReferenceId) SetForMarshal() {
	R.XMLName.Local = "t:ReferenceId"
}

func (R *ReferenceId) GetSchema() *Schema {
	return &SchemaTypes
}
