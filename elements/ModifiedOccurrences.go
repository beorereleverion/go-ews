package elements

// The ModifiedOccurrences element contains an array of recurring calendar item occurrences that have been modified so that they are different than the recurrence master item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/modifiedoccurrences
import "encoding/xml"

type ModifiedOccurrences struct {
	XMLName xml.Name

	// The Occurrence element represents a single modified occurrence of a recurring calendar item.
	Occurrence *Occurrence `xml:"Occurrence"`
}

func (M *ModifiedOccurrences) SetForMarshal() {
	M.XMLName.Local = "t:ModifiedOccurrences"
}

func (M *ModifiedOccurrences) GetSchema() *Schema {
	return &SchemaTypes
}
