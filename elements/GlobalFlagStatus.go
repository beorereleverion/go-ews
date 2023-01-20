package elements

// The GlobalFlagStatus element contains the aggregated flag status for all conversation items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalflagstatus
import "encoding/xml"

type GlobalFlagStatus struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Complete - Indicates the complete flag status.

	GlobalFlagStatusComplete string = `Complete`
	// Flagged - Indicates the flagged status.

	GlobalFlagStatusFlagged string = `Flagged`
	// NotFlagged - Indicates the not-flagged status.

	GlobalFlagStatusNotFlagged string = `NotFlagged`
)

func (G *GlobalFlagStatus) SetForMarshal() {
	G.XMLName.Local = "t:GlobalFlagStatus"
}

func (G *GlobalFlagStatus) GetSchema() *Schema {
	return &SchemaTypes
}
