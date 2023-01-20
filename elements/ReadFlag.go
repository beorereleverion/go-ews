package elements

// The ReadFlag element indicates the read state to set on items in a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/readflag
import "encoding/xml"

type ReadFlag struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ReadFlagfalse bool = false
	// true
	ReadFlagtrue bool = true
)

func (R *ReadFlag) SetForMarshal() {
	R.XMLName.Local = "m:ReadFlag"
}

func (R *ReadFlag) GetSchema() *Schema {
	return &SchemaMessages
}
