package elements

// The KeepProperties element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/keepproperties
import "encoding/xml"

type KeepProperties struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (K *KeepProperties) SetForMarshal() {
	K.XMLName.Local = "m:KeepProperties"
}

func (K *KeepProperties) GetSchema() *Schema {
	return &SchemaMessages
}
