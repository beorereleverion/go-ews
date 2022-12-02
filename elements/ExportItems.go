package elements

// The ExportItems element represents a request to export items from a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exportitems
type ExportItems struct {
	// The ItemIds element contains an array of item identifiers that identify the items to export from a mailbox.
	ItemIds *ItemIdsNonEmptyArrayOfItemIdsType `xml:"m:ItemIds"`
}
