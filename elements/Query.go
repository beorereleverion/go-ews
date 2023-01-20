package elements

// The Query element contains the search query for the hold.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/query
import "encoding/xml"

type Query struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (Q *Query) SetForMarshal() {
	Q.XMLName.Local = "m:Query"
}

func (Q *Query) GetSchema() *Schema {
	return &SchemaMessages
}
