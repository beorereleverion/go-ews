package elements

// The Location (ReminderMessageDataType) element specifies the location of the calendar item that the reminder is for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/location-remindermessagedatatype
import "encoding/xml"

type LocationReminderMessageDataType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (L *LocationReminderMessageDataType) SetForMarshal() {
	L.XMLName.Local = "t:Location"
}

func (L *LocationReminderMessageDataType) GetSchema() *Schema {
	return &SchemaTypes
}
