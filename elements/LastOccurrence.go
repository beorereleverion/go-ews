package elements

// The LastOccurrence element represents the last occurrence of a recurring calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastoccurrence
import "encoding/xml"

type LastOccurrence struct {
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

func (L *LastOccurrence) SetForMarshal() {
	L.XMLName.Local = "t:LastOccurrence"
}

func (L *LastOccurrence) GetSchema() *Schema {
	return &SchemaTypes
}
