package elements

// The PhysicalAddresses element contains a collection of physical addresses that are associated with a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/physicaladdresses
type PhysicalAddresses struct {
	// The Entry element describes a single physical address for a contact item.
	Entry *EntryPhysicalAddress `xml:"t:Entry"`
}
