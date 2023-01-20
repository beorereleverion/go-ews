package elements

// The NoEndRecurrence element describes the start date of an item recurrence pattern that does not have a defined end date.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/noendrecurrence
import "encoding/xml"

type NoEndRecurrence struct {
	XMLName xml.Name

	// The StartDate element represents the start date of a recurring task or calendar item.
	StartDate *StartDateRecurrence `xml:"StartDate"`
}

func (N *NoEndRecurrence) SetForMarshal() {
	N.XMLName.Local = "t:NoEndRecurrence"
}

func (N *NoEndRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
