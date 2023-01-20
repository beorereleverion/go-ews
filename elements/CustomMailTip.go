package elements

// The CustomMailTip element represents a customized mail tip message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/custommailtip
import "encoding/xml"

type CustomMailTip struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *CustomMailTip) SetForMarshal() {
	C.XMLName.Local = "t:CustomMailTip"
}

func (C *CustomMailTip) GetSchema() *Schema {
	return &SchemaTypes
}
