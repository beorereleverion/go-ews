package elements

// The FlagStatus element contains the aggregated flag status for conversation items in the current folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/flagstatus
import "encoding/xml"

type FlagStatus struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Complete - Indicates the complete flag status.

	FlagStatusComplete string = `Complete`
	// Flagged - Indicates the flagged status.

	FlagStatusFlagged string = `Flagged`
	// NotFlagged - Indicates the not-flagged status.

	FlagStatusNotFlagged string = `NotFlagged`
)

func (F *FlagStatus) SetForMarshal() {
	F.XMLName.Local = "t:FlagStatus"
}

func (F *FlagStatus) GetSchema() *Schema {
	return &SchemaTypes
}
