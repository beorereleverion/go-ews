package elements

// The DataType element describes the type of data that is shared by a shared folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datatype
import "encoding/xml"

type DataType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Calendar
	DataTypeCalendar string = `Calendar`
	// Contacts
	DataTypeContacts string = `Contacts`
)

func (D *DataType) SetForMarshal() {
	D.XMLName.Local = "m:DataType"
}

func (D *DataType) GetSchema() *Schema {
	return &SchemaMessages
}
