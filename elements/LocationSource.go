package elements

// The LocationSource element specifies information about the origin of the associated postal address, for example, a contact or a telephone book.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/locationsource
type LocationSource string

const (
	// Contact
	LocationSourceContact LocationSource = `Contact`
	// Device
	LocationSourceDevice LocationSource = `Device`
	// LocationServices
	LocationSourceLocationServices LocationSource = `LocationServices`
	// None
	LocationSourceNone LocationSource = `None`
	// PhonebookServices
	LocationSourcePhonebookServices LocationSource = `PhonebookServices`
	// Resource
	LocationSourceResource LocationSource = `Resource`
)
