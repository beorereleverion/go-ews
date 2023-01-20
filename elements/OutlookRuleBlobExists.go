package elements

// The OutlookRuleBlobExists element indicates whether a Microsoft Outlook rule blob exists in the user's mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/outlookruleblobexists
import "encoding/xml"

type OutlookRuleBlobExists struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	OutlookRuleBlobExistsfalse bool = false
	// true
	OutlookRuleBlobExiststrue bool = true
)

func (O *OutlookRuleBlobExists) SetForMarshal() {
	O.XMLName.Local = "m:OutlookRuleBlobExists"
}

func (O *OutlookRuleBlobExists) GetSchema() *Schema {
	return &SchemaMessages
}
