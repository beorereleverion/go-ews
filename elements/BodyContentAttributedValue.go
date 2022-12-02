package elements

// The BodyContentAttributedValue element specifies the body content of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bodycontentattributedvalue
type BodyContentAttributedValue struct {
	// The Attributions element specifies an array of attribution information for one or more of the contacts or Active Directory recipients aggregated into the associated persona.
	Attributions *AttributionsArrayOfPersonaAttributionsType `xml:"t:Attributions"`
	// The Value element specifies the value of a BodyContentAttributedValue element.
	Value *ValueBodyContentType `xml:"t:Value"`
}
