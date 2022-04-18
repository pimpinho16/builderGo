package builderGo

import (
	"encoding/xml"
)

type XMLJsonMessageBuilder struct {
	message Message
}

func (b *XMLJsonMessageBuilder) SetMessage(message string) MessageBuilder {
	b.message.Text = message
	return b
}

func (b *XMLJsonMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := xml.Marshal(b.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{Body: data, Format: "XML"}, nil
}

func (b *XMLJsonMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}
