package elements

// The GetImItemList element defines a request to get a list of instant messaging groups and contacts.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getimitemlist
type GetImItemList struct {
	// The ExtendedProperties element contains the extended properties used for the Unified Contact Store operations.
	ExtendedProperties *ExtendedPropertiesNonEmptyArrayOfExtendedFieldURIs `xml:"m:ExtendedProperties"`
}
