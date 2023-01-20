package elements

// The Size element specifies the total size of one or more mailbox items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/size-long
import "encoding/xml"

type Sizelong struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (S *Sizelong) SetForMarshal() {
	S.XMLName.Local = "t:Size"
}

func (S *Sizelong) GetSchema() *Schema {
	return &SchemaTypes
}
