package elements

// The Refiner element specifies a search refiner.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/refiner
type Refiner struct {
	// The Count element contains the number of conflicts in an UpdateItem operation response.
	Count *Count `xml:"t:Count"`
	// The Name element specifies a search refiner name.
	Name *Namestring `xml:"t:Name"`
	// The Token element contains a search refiner token.
	Token *TokenString `xml:"t:Token"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"t:Value"`
}
