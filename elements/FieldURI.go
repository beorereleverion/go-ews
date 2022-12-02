package elements

// The FieldURI element identifies frequently referenced properties by URI.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fielduri
type FieldURI struct {
	// Identifies the URI of the property.
	FieldURI *string `xml:"FieldURI,attr"`
}
