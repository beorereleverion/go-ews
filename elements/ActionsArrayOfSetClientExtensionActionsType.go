package elements

// The Actions element identifies an array of actions to perform on a client extension.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actions-arrayofsetclientextensionactionstype
import "encoding/xml"

type ActionsArrayOfSetClientExtensionActionsType struct {
	XMLName xml.Name

	// The Action element contains the action that the Exchange server should take on an app.
	Action *ActionSetClientExtensionActionType `xml:"Action"`
}

func (A *ActionsArrayOfSetClientExtensionActionsType) SetForMarshal() {
	A.XMLName.Local = "m:Actions"
}

func (A *ActionsArrayOfSetClientExtensionActionsType) GetSchema() *Schema {
	return &SchemaMessages
}
