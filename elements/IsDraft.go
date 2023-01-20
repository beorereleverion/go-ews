package elements

// The IsDraft element indicates whether an item has not yet been sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isdraft
import "encoding/xml"

type IsDraft struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsDraft) SetForMarshal() {
	I.XMLName.Local = "t:IsDraft"
}

func (I *IsDraft) GetSchema() *Schema {
	return &SchemaTypes
}
