package elements

// The Occurrence element represents a single modified occurrence of a recurring calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/occurrence
import "encoding/xml"

type Occurrence struct {
	XMLName xml.Name

	// The End element represents the end of a duration.
	End *End `xml:"End"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The OriginalStart element represents the original start time of a calendar item.
	OriginalStart *OriginalStart `xml:"OriginalStart"`
	// The Start element represents the start of a duration.
	Start *Start `xml:"Start"`
}

func (O *Occurrence) SetForMarshal() {
	O.XMLName.Local = "t:Occurrence"
}

func (O *Occurrence) GetSchema() *Schema {
	return &SchemaTypes
}
