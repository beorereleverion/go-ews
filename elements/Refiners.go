package elements

// The Refiners element specifies a list of one or more Refiner elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/refiners
import "encoding/xml"

type Refiners struct {
	XMLName xml.Name

	// The Refiner element specifies a search refiner.
	Refiner *Refiner `xml:"Refiner"`
}

func (R *Refiners) SetForMarshal() {
	R.XMLName.Local = "t:Refiners"
}

func (R *Refiners) GetSchema() *Schema {
	return &SchemaTypes
}
