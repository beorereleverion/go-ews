package elements

// The Refiners element specifies a list of one or more Refiner elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/refiners
type Refiners struct {
	// The Refiner element specifies a search refiner.
	Refiner *Refiner `xml:"t:Refiner"`
}
