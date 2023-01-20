package elements

// The HasChanged element indicates whether a user's photo has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/haschanged
import "encoding/xml"

type HasChanged struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// false
	HasChangedfalse string = `false`
	// true
	HasChangedtrue string = `true`
)

func (H *HasChanged) SetForMarshal() {
	H.XMLName.Local = "t:HasChanged"
}

func (H *HasChanged) GetSchema() *Schema {
	return &SchemaTypes
}
