package elements

// The InstallApp element specifies the request to install an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/installapp
import "encoding/xml"

type InstallApp struct {
	XMLName xml.Name

	// The Manifest element contains the base64-encoded app manifest file.
	Manifest *Manifest `xml:"Manifest"`
}

func (I *InstallApp) SetForMarshal() {
	I.XMLName.Local = "m:InstallApp"
}

func (I *InstallApp) GetSchema() *Schema {
	return &SchemaMessages
}
