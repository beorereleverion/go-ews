package elements

// The Errors element contains a property bag to store errors that are returned through the Web service.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/errors-ex15websvcsotherref
import "encoding/xml"

type Errors struct {
	XMLName xml.Name

	// The Properties element contains a list of one or more tracking properties.
	Properties *PropertiesArrayOfTrackingPropertiesType `xml:"Properties"`
}

func (E *Errors) SetForMarshal() {
	E.XMLName.Local = "m:Errors"
}

func (E *Errors) GetSchema() *Schema {
	return &SchemaMessages
}
