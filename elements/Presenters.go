package elements

// The Presenters element specifies the presenters for an online meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/presenters
import "encoding/xml"

type Presenters struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Disabled
	PresentersDisabled string = `Disabled`
	// Everyone
	PresentersEveryone string = `Everyone`
	// Internal
	PresentersInternal string = `Internal`
)

func (P *Presenters) SetForMarshal() {
	P.XMLName.Local = "t:Presenters"
}

func (P *Presenters) GetSchema() *Schema {
	return &SchemaTypes
}
