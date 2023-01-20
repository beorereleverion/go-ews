package elements

// The StoreEntryId element contains the Exchange store identifier of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/storeentryid
import "encoding/xml"

type StoreEntryId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}
