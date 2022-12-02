package elements

// The ActionType element indicates the type of action for the hold.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actiontype-holdactiontype
type ActionTypeHoldActionType string

const (
	// Create
	ActionTypeHoldActionTypeCreate ActionTypeHoldActionType = `Create`
	// Remove
	ActionTypeHoldActionTypeRemove ActionTypeHoldActionType = `Remove`
	// Update
	ActionTypeHoldActionTypeUpdate ActionTypeHoldActionType = `Update`
)
