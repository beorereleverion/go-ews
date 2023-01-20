package elements

// The DeletedOccurrenceStateDefinition specifies the state for a deleted occurrence of a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletedoccurrencestatedefinition
import "encoding/xml"

type DeletedOccurrenceStateDefinition struct {
	XMLName xml.Name

	// The IsOccurrencePresent element is intended for internal use only.
	IsOccurrencePresent *IsOccurrencePresent `xml:"IsOccurrencePresent"`
	// The Occurrence element represents the occurrence of the day of the week in the month that the time zone transition occurs.
	Occurrence *OccurrenceTimeZoneTransition `xml:"Occurrence"`
}

func (D *DeletedOccurrenceStateDefinition) SetForMarshal() {
	D.XMLName.Local = "t:DeletedOccurrenceStateDefinition"
}

func (D *DeletedOccurrenceStateDefinition) GetSchema() *Schema {
	return &SchemaTypes
}
