package elements

// The LobbyBypass element specifies the online meeting setting to bypass the virtual lobby.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lobbybypass
type LobbyBypass string

const (
	// Disabled
	LobbyBypassDisabled LobbyBypass = `Disabled`
	// EnabledForGatewayParticipants
	LobbyBypassEnabledForGatewayParticipants LobbyBypass = `EnabledForGatewayParticipants`
)
