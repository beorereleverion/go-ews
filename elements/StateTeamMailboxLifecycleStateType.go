package elements

// The State element contains the lifecycle state that is set on a site mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/state-teammailboxlifecyclestatetype
import "encoding/xml"

type StateTeamMailboxLifecycleStateType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Active
	StateTeamMailboxLifecycleStateTypeActive string = `Active`
	// Closed
	StateTeamMailboxLifecycleStateTypeClosed string = `Closed`
	// PendingDelete
	StateTeamMailboxLifecycleStateTypePendingDelete string = `PendingDelete`
	// Unlinked
	StateTeamMailboxLifecycleStateTypeUnlinked string = `Unlinked`
)

func (S *StateTeamMailboxLifecycleStateType) SetForMarshal() {
	S.XMLName.Local = "m:State"
}

func (S *StateTeamMailboxLifecycleStateType) GetSchema() *Schema {
	return &SchemaMessages
}
