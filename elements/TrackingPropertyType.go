package elements

// The TrackingPropertyType element represents a name and value pair of strings that is used to create properties for message tracking reports.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/trackingpropertytype
import "encoding/xml"

type TrackingPropertyType struct {
	XMLName xml.Name

	// The Name element represents the property name for a message tracking report.
	Name *NameMessageTracking `xml:"Name"`
	// The Value element represents the property value for a message tracking report.
	Value *ValueMessageTracking `xml:"Value"`
}

func (T *TrackingPropertyType) SetForMarshal() {
	T.XMLName.Local = "m:TrackingPropertyType"
}

func (T *TrackingPropertyType) GetSchema() *Schema {
	return &SchemaMessages
}
