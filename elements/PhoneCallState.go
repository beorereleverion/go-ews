package elements

// The PhoneCallState element specifies the current state for a phone call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonecallstate
type PhoneCallState string

const (
	// Alerted
	PhoneCallStateAlerted PhoneCallState = `Alerted`
	// Connected
	PhoneCallStateConnected PhoneCallState = `Connected`
	// Connecting
	PhoneCallStateConnecting PhoneCallState = `Connecting`
	// Disconnected
	PhoneCallStateDisconnected PhoneCallState = `Disconnected`
	// Forwarding
	PhoneCallStateForwarding PhoneCallState = `Forwarding`
	// Idle
	PhoneCallStateIdle PhoneCallState = `Idle`
	// Incoming
	PhoneCallStateIncoming PhoneCallState = `Incoming`
	// Transferring
	PhoneCallStateTransferring PhoneCallState = `Transferring`
)
