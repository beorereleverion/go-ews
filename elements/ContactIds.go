package elements

// The ContactIds element contains an array of contact item identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contactids
import "encoding/xml"

type ContactIds struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The OccurrenceItemId element identifies a single occurrence of a recurring item.
	OccurrenceItemId *OccurrenceItemId `xml:"OccurrenceItemId"`
	// The RecurringMasterItemId element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
	RecurringMasterItemId *RecurringMasterItemId `xml:"RecurringMasterItemId"`
	// The RecurringMasterItemIdRanges element specifies an array of occurrence ranges.
	RecurringMasterItemIdRanges *RecurringMasterItemIdRanges `xml:"RecurringMasterItemIdRanges"`
}

func (C *ContactIds) SetForMarshal() {
	C.XMLName.Local = "t:ContactIds"
}

func (C *ContactIds) GetSchema() *Schema {
	return &SchemaTypes
}
