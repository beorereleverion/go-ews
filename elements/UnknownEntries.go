package elements

// The UnknownEntries element contains an array of unknown permission entries that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unknownentries
type UnknownEntries struct {
	// The UnknownEntry element represents a single unknown permission entry that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	UnknownEntry *UnknownEntry `xml:"t:UnknownEntry"`
}
