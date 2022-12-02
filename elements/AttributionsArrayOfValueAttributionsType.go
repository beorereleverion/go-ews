package elements

// The Attributions element specifies an array of attributions for its associated Value element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attributions-arrayofvalueattributionstype
type AttributionsArrayOfValueAttributionsType struct {
	// The Attribution element specifies a string used to identify an attribute of a persona.
	Attribution *Attributionstring `xml:"t:Attribution"`
}
