package elements

// The Accuracy element specifies the accuracy of the latitude and longitude of the associated postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/accuracy
import "encoding/xml"

type Accuracy struct {
	XMLName xml.Name
	TEXT    float64 `xml:",chardata"`
}

func (A *Accuracy) SetForMarshal() {
	A.XMLName.Local = "t:Accuracy"
}

func (A *Accuracy) GetSchema() *Schema {
	return &SchemaTypes
}
