package elements

// The AltitudeAccuracy element specifies the accuracy of the altitude property for a postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/altitudeaccuracy
import "encoding/xml"

type AltitudeAccuracy struct {
	XMLName xml.Name
	TEXT    float64 `xml:",chardata"`
}

func (A *AltitudeAccuracy) SetForMarshal() {
	A.XMLName.Local = "t:AltitudeAccuracy"
}

func (A *AltitudeAccuracy) GetSchema() *Schema {
	return &SchemaTypes
}
