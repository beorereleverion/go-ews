package elements

// The ConnectionFailureCause element specifies the reason for a disconnection from a telephone call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/connectionfailurecause
import "encoding/xml"

type ConnectionFailureCause struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// NoAnswer
	ConnectionFailureCauseNoAnswer string = `NoAnswer`
	// None
	ConnectionFailureCauseNone string = `None`
	// Other
	ConnectionFailureCauseOther string = `Other`
	// Unavailable
	ConnectionFailureCauseUnavailable string = `Unavailable`
	// UserBusy
	ConnectionFailureCauseUserBusy string = `UserBusy`
)

func (C *ConnectionFailureCause) SetForMarshal() {
	C.XMLName.Local = "t:ConnectionFailureCause"
}

func (C *ConnectionFailureCause) GetSchema() *Schema {
	return &SchemaTypes
}
