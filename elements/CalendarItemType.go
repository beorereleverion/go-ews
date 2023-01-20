package elements

// The CalendarItemType element represents the type of a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendaritemtype
import "encoding/xml"

type CalendarItemType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Exception The item is an exception to a recurring calendar item.

	CalendarItemTypeException string = `Exception`
	// Occurrence The item is an occurrence of a recurring calendar item.

	CalendarItemTypeOccurrence string = `Occurrence`
	// RecurringMaster The item is master for a set of recurring calendar items.

	CalendarItemTypeRecurringMaster string = `RecurringMaster`
	// Single The item is not associated with a recurring calendar item.

	CalendarItemTypeSingle string = `Single`
)

func (C *CalendarItemType) SetForMarshal() {
	C.XMLName.Local = "t:CalendarItemType"
}

func (C *CalendarItemType) GetSchema() *Schema {
	return &SchemaTypes
}
