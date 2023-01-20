package elements

// The DisableApp element specifies a request to disable an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/disableapp
import "encoding/xml"

type DisableApp struct {
	XMLName xml.Name

	// The DisableReason element specifies the reason for disabling an app.
	DisableReason *DisableReason `xml:"DisableReason"`
	// The ID element specifies the identifier of an app.
	ID *IDString `xml:"ID"`
}

func (D *DisableApp) SetForMarshal() {
	D.XMLName.Local = "m:DisableApp"
}

func (D *DisableApp) GetSchema() *Schema {
	return &SchemaMessages
}
