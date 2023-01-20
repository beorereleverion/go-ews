package elements

// The Data element contains the data of a single exported item or an item to upload into a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/data-base64binary
import "encoding/xml"

type Database64Binary struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *Database64Binary) SetForMarshal() {
	D.XMLName.Local = "m:Data"
}

func (D *Database64Binary) GetSchema() *Schema {
	return &SchemaMessages
}
