package elements

// The UninstallApp element specifies a request to uninstall an app by its identifier.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uninstallapp
import "encoding/xml"

type UninstallApp struct {
	XMLName xml.Name

	// The ID element specifies the identifier of an app.
	ID *IDString `xml:"ID"`
}

func (U *UninstallApp) SetForMarshal() {
	U.XMLName.Local = "t:UninstallApp"
}

func (U *UninstallApp) GetSchema() *Schema {
	return &SchemaTypes
}
