package elements

// The Refiner element specifies a search refiner.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/refiner
import "encoding/xml"

type Refiner struct {
	XMLName xml.Name

	// The Count element contains the number of conflicts in an UpdateItem operation response.
	Count *Count `xml:"Count"`
	// The Name element specifies a search refiner name.
	Name *Namestring `xml:"Name"`
	// The Token element contains a search refiner token.
	Token *TokenString `xml:"Token"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"Value"`
}

func (R *Refiner) SetForMarshal() {
	R.XMLName.Local = "t:Refiner"
}

func (R *Refiner) GetSchema() *Schema {
	return &SchemaTypes
}
