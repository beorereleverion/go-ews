package elements

// The Size element represents the size in bytes of an item or all the items in a conversation in the current folder. This property is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/size
import "encoding/xml"

type Size struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (S *Size) SetForMarshal() {
	S.XMLName.Local = "t:Size"
}

func (S *Size) GetSchema() *Schema {
	return &SchemaTypes
}
