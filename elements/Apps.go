package elements

// The Apps element contains information about all the XML manifest files for apps installed in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/apps
import "encoding/xml"

type Apps struct {
	XMLName xml.Name

	// The App element contains information about an XML manifest file for a mail app that is installed in a mailbox.
	App *App `xml:"App"`
}

func (A *Apps) SetForMarshal() {
	A.XMLName.Local = "m:Apps"
}

func (A *Apps) GetSchema() *Schema {
	return &SchemaMessages
}
