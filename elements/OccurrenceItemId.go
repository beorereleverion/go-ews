package elements

// The OccurrenceItemId element identifies a single occurrence of a recurring item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/occurrenceitemid
import "encoding/xml"

type OccurrenceItemId struct {
	XMLName xml.Name

	// Identifies a specific version of the recurring master or an item occurrence. If either the recurring master or any of its occurrences change, the ChangeKey changes. The ChangeKey is the same for the recurring master and all occurrences.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies the index of the item occurrence. This attribute is required. This value represents an integer.
	InstanceIndex *string `xml:"InstanceIndex,attr"`
	// Identifies the recurring master of a recurring item. This attribute is required.
	RecurringMasterId *string `xml:"RecurringMasterId,attr"`
}

func (O *OccurrenceItemId) SetForMarshal() {
	O.XMLName.Local = "t:OccurrenceItemId"
}

func (O *OccurrenceItemId) GetSchema() *Schema {
	return &SchemaTypes
}
