package elements

// The AppointmentSequenceNumber element specifies the sequence number of a version of an appointment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/appointmentsequencenumber
import "encoding/xml"

type AppointmentSequenceNumber struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (A *AppointmentSequenceNumber) SetForMarshal() {
	A.XMLName.Local = "t:AppointmentSequenceNumber"
}

func (A *AppointmentSequenceNumber) GetSchema() *Schema {
	return &SchemaTypes
}
