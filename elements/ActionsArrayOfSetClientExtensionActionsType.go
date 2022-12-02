package elements

// The Actions element identifies an array of actions to perform on a client extension.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actions-arrayofsetclientextensionactionstype
type ActionsArrayOfSetClientExtensionActionsType struct {
	// The Action element contains the action that the Exchange server should take on an app.
	Action *ActionSetClientExtensionActionType `xml:"t:Action"`
}
