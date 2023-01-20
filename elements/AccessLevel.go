package elements

// The AccessLevel element specifies the access level for an online meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/accesslevel
import "encoding/xml"

type AccessLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// The access level is open to all.
	AccessLevelEveryone string = `Everyone`
	// The access level is internal only.
	AccessLevelInternal string = `Internal`
	// The access level is invited participants only.
	AccessLevelInvited string = `Invited`
	// The access level is locked.
	AccessLevelLocked string = `Locked`
)

func (A *AccessLevel) SetForMarshal() {
	A.XMLName.Local = "t:AccessLevel"
}

func (A *AccessLevel) GetSchema() *Schema {
	return &SchemaTypes
}
