package elements

// The IsCancelled element indicates whether an appointment or meeting has been canceled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/iscancelled
import "encoding/xml"

type IsCancelled struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsCancelled) SetForMarshal() {
	I.XMLName.Local = "t:IsCancelled"
}

func (I *IsCancelled) GetSchema() *Schema {
	return &SchemaTypes
}
