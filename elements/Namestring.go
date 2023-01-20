package elements

// The Name element specifies a search refiner name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/name-string
import "encoding/xml"

type Namestring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *Namestring) SetForMarshal() {
	N.XMLName.Local = "t:Name"
}

func (N *Namestring) GetSchema() *Schema {
	return &SchemaTypes
}
