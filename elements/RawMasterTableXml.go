package elements

// The RawMasterTableXml element is not used.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rawmastertablexml
import "encoding/xml"

type RawMasterTableXml struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *RawMasterTableXml) SetForMarshal() {
	R.XMLName.Local = "m:RawMasterTableXml"
}

func (R *RawMasterTableXml) GetSchema() *Schema {
	return &SchemaMessages
}
