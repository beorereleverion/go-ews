package elements

// The RecurringMasterItemIdRanges element specifies an array of occurrence ranges.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurringmasteritemidranges
type RecurringMasterItemIdRanges struct {
	// The Ranges element specifies an array of recurrence ranges.
	Ranges *Ranges `xml:"t:Ranges"`
	// The text value of the ChangeKey attribute is the recurring master item's change key. This is a string value.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The text value of the Id attribute is a recurring master item's unique identifier. This is a string value.
	Id *string `xml:"Id,attr"`
}
