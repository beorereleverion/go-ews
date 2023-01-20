package elements

// The DeletedOccurrence element represents a deleted occurrence of a recurring calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletedoccurrence
import "encoding/xml"

type DeletedOccurrence struct {
	XMLName xml.Name

	// The Start element represents the start of a duration.
	Start *Start `xml:"Start"`
}

func (D *DeletedOccurrence) SetForMarshal() {
	D.XMLName.Local = "t:DeletedOccurrence"
}

func (D *DeletedOccurrence) GetSchema() *Schema {
	return &SchemaTypes
}
