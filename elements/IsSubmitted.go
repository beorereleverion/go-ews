package elements

// The IsSubmitted element indicates whether an item has been submitted to the Outbox default folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/issubmitted
import "encoding/xml"

type IsSubmitted struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsSubmitted) SetForMarshal() {
	I.XMLName.Local = "t:IsSubmitted"
}

func (I *IsSubmitted) GetSchema() *Schema {
	return &SchemaTypes
}
