package elements

// The Attributions element specifies an array of attribution information for one or more of the contacts or Active Directory recipients aggregated into the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attributions-arrayofpersonaattributionstype
type AttributionsArrayOfPersonaAttributionsType struct {
	// The Attribution element specifies an instance in an array of attributes for a PersonaType element.
	Attribution *AttributionPersonaAttributionType `xml:"t:Attribution"`
}
