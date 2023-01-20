package elements

// The EndDateRecurrence element describes the start date and the end date of an item recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/enddaterecurrence
import "encoding/xml"

type EndDateRecurrence struct {
	XMLName xml.Name

	// The EndDate element represents the end date of a recurring task or a calendar item that has the EndDateRecurrence pattern type.
	EndDate *EndDateRecurrence `xml:"EndDate"`
	// The StartDate element represents the start date of a recurring task or calendar item.
	StartDate *StartDateRecurrence `xml:"StartDate"`
}

func (E *EndDateRecurrence) SetForMarshal() {
	E.XMLName.Local = "t:EndDateRecurrence"
}

func (E *EndDateRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
