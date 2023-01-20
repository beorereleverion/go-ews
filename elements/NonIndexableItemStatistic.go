package elements

// The NonIndexableItemStatistic element contains a single statistic for an item that could not be indexed
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nonindexableitemstatistic
import "encoding/xml"

type NonIndexableItemStatistic struct {
	XMLName xml.Name

	// The ErrorMessage element represents the reason for the validation error.
	ErrorMessage *ErrorMessage `xml:"ErrorMessage"`
	// The ItemCount element specifies the total number of items in a search result.
	ItemCount *ItemCount `xml:"ItemCount"`
	// The Mailbox element contains an identifier for a mailbox.
	Mailbox *Mailboxstring `xml:"Mailbox"`
}

func (N *NonIndexableItemStatistic) SetForMarshal() {
	N.XMLName.Local = "m:NonIndexableItemStatistic"
}

func (N *NonIndexableItemStatistic) GetSchema() *Schema {
	return &SchemaMessages
}
