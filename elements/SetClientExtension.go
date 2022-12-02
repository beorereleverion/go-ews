package elements

// The SetClientExtension element contains a request to set a client extension.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setclientextension
type SetClientExtension struct {
	// The Actions element identifies an array of actions to perform on a client extension.
	Actions *ActionsArrayOfSetClientExtensionActionsType `xml:"m:Actions"`
}
