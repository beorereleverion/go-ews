package elements

// The OnlineMeetingSettings element specifies the settings for online meetings.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/onlinemeetingsettings
type OnlineMeetingSettings struct {
	// The AccessLevel element specifies the access level for an online meeting.
	AccessLevel *AccessLevel `xml:"t:AccessLevel"`
	// The LobbyBypass element specifies the online meeting setting to bypass the virtual lobby.
	LobbyBypass *LobbyBypass `xml:"LobbyBypass"`
	// The Presenters element specifies the presenters for an online meeting.
	Presenters *Presenters `xml:"t:Presenters"`
}
