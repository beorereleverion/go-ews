package elements

// The ActionType element indicates the type of action for the hold.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actiontype-holdactiontype
import "encoding/xml"

type ActionTypeHoldActionType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Create
	ActionTypeHoldActionTypeCreate string = `Create`
	// Remove
	ActionTypeHoldActionTypeRemove string = `Remove`
	// Update
	ActionTypeHoldActionTypeUpdate string = `Update`
)

func (A *ActionTypeHoldActionType) SetForMarshal() {
	A.XMLName.Local = "t:ActionType"
}

func (A *ActionTypeHoldActionType) GetSchema() *Schema {
	return &SchemaTypes
}
