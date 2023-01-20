package elements

// The GetNonIndexableItemStatistics element specifies a request to retrieve nonindexable item statistics.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getnonindexableitemstatistics
import "encoding/xml"

type GetNonIndexableItemStatistics struct {
	XMLName xml.Name

	// The Mailboxes element specifies an array of mailboxes identified by legacy distinguished name.
	Mailboxes *MailboxesNonEmptyArrayOfLegacyDNsType `xml:"Mailboxes"`
}

func (G *GetNonIndexableItemStatistics) SetForMarshal() {
	G.XMLName.Local = "m:GetNonIndexableItemStatistics"
}

func (G *GetNonIndexableItemStatistics) GetSchema() *Schema {
	return &SchemaMessages
}
