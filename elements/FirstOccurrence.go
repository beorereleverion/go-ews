package elements

// The FirstOccurrence element represents the first occurrence of a recurring calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/firstoccurrence
type FirstOccurrence struct {
	// The End element represents the end of a duration.
	End *End `xml:"t:End"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The OriginalStart element represents the original start time of a calendar item.
	OriginalStart *OriginalStart `xml:"t:OriginalStart"`
	// The Start element represents the start of a duration.
	Start *Start `xml:"t:Start"`
}
