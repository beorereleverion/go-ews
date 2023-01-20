package elements

// The WorkCity element specifies the city where the associated persona works.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workcity
import "encoding/xml"

type WorkCity struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (W *WorkCity) SetForMarshal() {
	W.XMLName.Local = "t:WorkCity"
}

func (W *WorkCity) GetSchema() *Schema {
	return &SchemaTypes
}
