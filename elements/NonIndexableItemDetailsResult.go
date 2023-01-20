package elements

// The NonIndexableItemDetailsResult element specifies the results of the GetNonIndexableItemDetails WSDL operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nonindexableitemdetailsresult
import "encoding/xml"

type NonIndexableItemDetailsResult struct {
	XMLName xml.Name

	// The FailedMailboxes element specifies an array of mailboxes that failed on search.
	FailedMailboxes *FailedMailboxes `xml:"FailedMailboxes"`
	// The Items element contains an array of item details for non-indexable items.
	Items *ItemsArrayOfNonIndexableItemDetailsType `xml:"Items"`
}

func (N *NonIndexableItemDetailsResult) SetForMarshal() {
	N.XMLName.Local = "m:NonIndexableItemDetailsResult"
}

func (N *NonIndexableItemDetailsResult) GetSchema() *Schema {
	return &SchemaMessages
}
