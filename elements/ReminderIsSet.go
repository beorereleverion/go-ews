package elements

// The ReminderIsSet element indicates whether a reminder has been set for an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminderisset
import "encoding/xml"

type ReminderIsSet struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (R *ReminderIsSet) SetForMarshal() {
	R.XMLName.Local = "t:ReminderIsSet"
}

func (R *ReminderIsSet) GetSchema() *Schema {
	return &SchemaTypes
}
