package elements

// The Mileage element represents mileage for a task or contact item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mileage
import "encoding/xml"

type Mileage struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *Mileage) SetForMarshal() {
	M.XMLName.Local = "t:Mileage"
}

func (M *Mileage) GetSchema() *Schema {
	return &SchemaTypes
}
