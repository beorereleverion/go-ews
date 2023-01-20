package elements

// The LocationSource element specifies information about the origin of the associated postal address, for example, a contact or a telephone book.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/locationsource
import "encoding/xml"

type LocationSource struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Contact
	LocationSourceContact string = `Contact`
	// Device
	LocationSourceDevice string = `Device`
	// LocationServices
	LocationSourceLocationServices string = `LocationServices`
	// None
	LocationSourceNone string = `None`
	// PhonebookServices
	LocationSourcePhonebookServices string = `PhonebookServices`
	// Resource
	LocationSourceResource string = `Resource`
)
