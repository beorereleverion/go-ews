package elements

// The Metadata element contains metadata about the mail app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/metadata-ex15websvcsotherref
import "encoding/xml"

type Metadata struct {
	XMLName xml.Name

	// The ActionUrl element identifies the URL that the user should navigate to, in order to fix an issue indicated by the AppStatus element.
	ActionUrl *ActionUrl `xml:"ActionUrl"`
	// The AppStatus element value indicates the status of the mail app.
	AppStatus *AppStatus `xml:"AppStatus"`
	// The EndNodeUrl element specifies the URL for the mail app in the Office Store.
	EndNodeUrl *EndNodeUrl `xml:"EndNodeUrl"`
}

func (M *Metadata) SetForMarshal() {
	M.XMLName.Local = "t:Metadata"
}

func (M *Metadata) GetSchema() *Schema {
	return &SchemaTypes
}
