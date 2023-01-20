package elements

// The PermanentDelete element indicates whether messages are to be permanently deleted and not saved to the Deleted Items folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/permanentdelete
import "encoding/xml"

type PermanentDelete struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	PermanentDeletefalse bool = false
	// true
	PermanentDeletetrue bool = true
)

func (P *PermanentDelete) SetForMarshal() {
	P.XMLName.Local = "m:PermanentDelete"
}

func (P *PermanentDelete) GetSchema() *Schema {
	return &SchemaMessages
}
