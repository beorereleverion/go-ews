package elements

// The Count element contains the number of conflicts in an UpdateItem operation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/count
import "encoding/xml"

type Count struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (C *Count) SetForMarshal() {
	C.XMLName.Local = "t:Count"
}

func (C *Count) GetSchema() *Schema {
	return &SchemaTypes
}
