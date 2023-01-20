package elements

// The IsRecurring element indicates whether a calendar item, meeting request, or task is part of a recurring item. This element is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isrecurring
import "encoding/xml"

type IsRecurring struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsRecurring) SetForMarshal() {
	I.XMLName.Local = "t:IsRecurring"
}

func (I *IsRecurring) GetSchema() *Schema {
	return &SchemaTypes
}
