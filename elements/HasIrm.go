package elements

// The HasIrm element specifies whether at least one message in the conversation and the current folder is an IRM protected message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasirm
import "encoding/xml"

type HasIrm struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	HasIrmfalse bool = false
	// true
	HasIrmtrue bool = true
)

func (H *HasIrm) SetForMarshal() {
	H.XMLName.Local = "t:HasIrm"
}

func (H *HasIrm) GetSchema() *Schema {
	return &SchemaTypes
}
