package elements

// The Position element specifies the position of an entity extracted from a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/position
import "encoding/xml"

type Position struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// LatestReply - the extracted entity originates from the latest reply to the message.

	PositionLatestReply string = `LatestReply`
	// Other - the extracted entity originates from an undefined part of the message.

	PositionOther string = `Other`
	// Signature - the extracted entity originates from the message signature.

	PositionSignature string = `Signature`
	// Subject - the extracted entity originates from the message subject.

	PositionSubject string = `Subject`
)

func (P *Position) SetForMarshal() {
	P.XMLName.Local = "t:Position"
}

func (P *Position) GetSchema() *Schema {
	return &SchemaTypes
}
