package elements

// The SendPrompt element specifies the type of action allowed for a voting option.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendprompt
import "encoding/xml"

type SendPrompt struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// None
	SendPromptNone string = `None`
	// Send
	SendPromptSend string = `Send`
	// VotingOption
	SendPromptVotingOption string = `VotingOption`
)

func (S *SendPrompt) SetForMarshal() {
	S.XMLName.Local = "t:SendPrompt"
}

func (S *SendPrompt) GetSchema() *Schema {
	return &SchemaTypes
}
