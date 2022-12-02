package elements

// The AddressListId element specifies the identifier of an address list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addresslistid
type AddressListId struct {
	// A string address list identifier. This attribute is required.
	Id *string `xml:"Id,attr"`
}
