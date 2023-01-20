package elements

// The SetClientExtension element contains a request to set a client extension.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setclientextension
import "encoding/xml"

type SetClientExtension struct {
	XMLName xml.Name

	// The Actions element identifies an array of actions to perform on a client extension.
	Actions *ActionsArrayOfSetClientExtensionActionsType `xml:"Actions"`
}

func (S *SetClientExtension) SetForMarshal() {
	S.XMLName.Local = "m:SetClientExtension"
}

func (S *SetClientExtension) GetSchema() *Schema {
	return &SchemaMessages
}
