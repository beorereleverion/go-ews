package elements

// The OnlineMeetingSettings element specifies the settings for online meetings.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/onlinemeetingsettings
import "encoding/xml"

type OnlineMeetingSettings struct {
	XMLName xml.Name

	// The AccessLevel element specifies the access level for an online meeting.
	AccessLevel *AccessLevel `xml:"AccessLevel"`
	// The LobbyBypass element specifies the online meeting setting to bypass the virtual lobby.
	LobbyBypass *LobbyBypass `xml:"LobbyBypass"`
	// The Presenters element specifies the presenters for an online meeting.
	Presenters *Presenters `xml:"Presenters"`
}

func (O *OnlineMeetingSettings) SetForMarshal() {
	O.XMLName.Local = "t:OnlineMeetingSettings"
}

func (O *OnlineMeetingSettings) GetSchema() *Schema {
	return &SchemaTypes
}
