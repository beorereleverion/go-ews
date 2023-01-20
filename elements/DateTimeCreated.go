package elements

// The DateTimeCreated element represents the date and time that an item in the mailbox was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimecreated
import "time"

import "encoding/xml"

type DateTimeCreated struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (D *DateTimeCreated) SetForMarshal() {
	D.XMLName.Local = "t:DateTimeCreated"
}

func (D *DateTimeCreated) GetSchema() *Schema {
	return &SchemaTypes
}
