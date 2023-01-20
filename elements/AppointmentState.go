package elements

// The AppointmentState element specifies the status of the appointment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/appointmentstate
import "encoding/xml"

type AppointmentState struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (A *AppointmentState) SetForMarshal() {
	A.XMLName.Local = "t:AppointmentState"
}

func (A *AppointmentState) GetSchema() *Schema {
	return &SchemaTypes
}
