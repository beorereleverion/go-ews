package elements

// The PhoneCallState element specifies the current state for a phone call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonecallstate
import "encoding/xml"

type PhoneCallState struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Alerted
	PhoneCallStateAlerted string = `Alerted`
	// Connected
	PhoneCallStateConnected string = `Connected`
	// Connecting
	PhoneCallStateConnecting string = `Connecting`
	// Disconnected
	PhoneCallStateDisconnected string = `Disconnected`
	// Forwarding
	PhoneCallStateForwarding string = `Forwarding`
	// Idle
	PhoneCallStateIdle string = `Idle`
	// Incoming
	PhoneCallStateIncoming string = `Incoming`
	// Transferring
	PhoneCallStateTransferring string = `Transferring`
)

func (P *PhoneCallState) SetForMarshal() {
	P.XMLName.Local = "t:PhoneCallState"
}

func (P *PhoneCallState) GetSchema() *Schema {
	return &SchemaTypes
}
