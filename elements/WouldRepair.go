package elements

// The WouldRepair element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/wouldrepair
import "encoding/xml"

type WouldRepair struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (W *WouldRepair) SetForMarshal() {
	W.XMLName.Local = "t:WouldRepair"
}

func (W *WouldRepair) GetSchema() *Schema {
	return &SchemaTypes
}
