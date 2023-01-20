package elements

// The UserOptions element specifies the list of voting options on a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/useroptions
import "encoding/xml"

type UserOptions struct {
	XMLName xml.Name

	// The VotingOptionData element specifies information about each voting option.
	VotingOptionData *VotingOptionData `xml:"VotingOptionData"`
}

func (U *UserOptions) SetForMarshal() {
	U.XMLName.Local = "t:UserOptions"
}

func (U *UserOptions) GetSchema() *Schema {
	return &SchemaTypes
}
