package elements

// The EndTimeInMinutes element represents the end of the working day for a mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/endtimeinminutes
import "encoding/xml"

type EndTimeInMinutes struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (E *EndTimeInMinutes) SetForMarshal() {
	E.XMLName.Local = "t:EndTimeInMinutes"
}

func (E *EndTimeInMinutes) GetSchema() *Schema {
	return &SchemaTypes
}
