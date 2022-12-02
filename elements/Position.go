package elements

// The Position element specifies the position of an entity extracted from a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/position
type Position string

const (
	// LatestReply - the extracted entity originates from the latest reply to the message.

	PositionLatestReply Position = `LatestReply`
	// Other - the extracted entity originates from an undefined part of the message.

	PositionOther Position = `Other`
	// Signature - the extracted entity originates from the message signature.

	PositionSignature Position = `Signature`
	// Subject - the extracted entity originates from the message subject.

	PositionSubject Position = `Subject`
)
