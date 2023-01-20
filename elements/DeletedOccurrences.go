package elements

// The DeletedOccurrences element contains an array of deleted occurrences of a recurring calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletedoccurrences
import "encoding/xml"

type DeletedOccurrences struct {
	XMLName xml.Name

	// The DeletedOccurrence element represents a deleted occurrence of a recurring calendar item.
	DeletedOccurrence *DeletedOccurrence `xml:"DeletedOccurrence"`
}

func (D *DeletedOccurrences) SetForMarshal() {
	D.XMLName.Local = "t:DeletedOccurrences"
}

func (D *DeletedOccurrences) GetSchema() *Schema {
	return &SchemaTypes
}
