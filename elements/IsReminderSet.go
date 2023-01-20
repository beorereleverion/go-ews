package elements

// The IsReminderSet element indicates whether a reminder has been set for the calendar event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isreminderset
import "encoding/xml"

type IsReminderSet struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsReminderSetfalse bool = false
	// true
	IsReminderSettrue bool = true
)

func (I *IsReminderSet) SetForMarshal() {
	I.XMLName.Local = "t:IsReminderSet"
}

func (I *IsReminderSet) GetSchema() *Schema {
	return &SchemaTypes
}
