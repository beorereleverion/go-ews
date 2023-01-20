package elements

// The NumberedRecurrence element describes the start date and the number of occurrences of a recurring item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/numberedrecurrence
import "encoding/xml"

type NumberedRecurrence struct {
	XMLName xml.Name

	// The NumberOfOccurrences element contains the number of occurrences of a recurring item.
	NumberOfOccurrences *NumberOfOccurrences `xml:"NumberOfOccurrences"`
	// The StartDate element represents the start date of a recurring task or calendar item.
	StartDate *StartDateRecurrence `xml:"StartDate"`
}

func (N *NumberedRecurrence) SetForMarshal() {
	N.XMLName.Local = "t:NumberedRecurrence"
}

func (N *NumberedRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
