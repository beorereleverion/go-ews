package elements

// The Year element is used to define a time zone that changes depending on the year. This element is optional. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/year
import "encoding/xml"

type Year struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (Y *Year) SetForMarshal() {
	Y.XMLName.Local = "t:Year"
}

func (Y *Year) GetSchema() *Schema {
	return &SchemaTypes
}
