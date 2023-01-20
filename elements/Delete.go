package elements

// The Delete element indicates whether a client can delete a folder or item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delete
import "encoding/xml"

type Delete struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	Deletefalse bool = false
	// true
	Deletetrue bool = true
)

func (D *Delete) SetForMarshal() {
	D.XMLName.Local = "t:Delete"
}

func (D *Delete) GetSchema() *Schema {
	return &SchemaTypes
}
