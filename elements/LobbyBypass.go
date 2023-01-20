package elements

// The LobbyBypass element specifies the online meeting setting to bypass the virtual lobby.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lobbybypass
import "encoding/xml"

type LobbyBypass struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Disabled
	LobbyBypassDisabled string = `Disabled`
	// EnabledForGatewayParticipants
	LobbyBypassEnabledForGatewayParticipants string = `EnabledForGatewayParticipants`
)
