package elements

// The App element contains information about an XML manifest file for a mail app that is installed in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/app
import "encoding/xml"

type App struct {
	XMLName xml.Name

	// The Manifest element contains the base64-encoded app manifest file.
	Manifest *Manifest `xml:"Manifest"`
	// The Metadata element contains metadata about the mail app.
	Metadata *Metadata `xml:"Metadata"`
}

func (A *App) SetForMarshal() {
	A.XMLName.Local = "t:App"
}

func (A *App) GetSchema() *Schema {
	return &SchemaTypes
}
