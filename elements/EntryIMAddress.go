package elements

// The Entry element represents an instant messaging (IM) address for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/entry-imaddress
type EntryIMAddress struct {
	// Identifies the IM address.The following are the possible values for this attribute:- ImAddress1  - ImAddress2  - ImAddress3
	Key *string `xml:"Key,attr"`
}
