package elements

// The StartTimeInMinutes element represents the start of the working day for a mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/starttimeinminutes
import "encoding/xml"

type StartTimeInMinutes struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (S *StartTimeInMinutes) SetForMarshal() {
	S.XMLName.Local = "t:StartTimeInMinutes"
}

func (S *StartTimeInMinutes) GetSchema() *Schema {
	return &SchemaTypes
}
