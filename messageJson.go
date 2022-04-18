package builderGo

import "encoding/json"

type JSONMessageBuilder struct {
	message Message
}

func (b *JSONMessageBuilder) SetMessage(message string) MessageBuilder {
	b.message.Text = message
	return b
}

func (b *JSONMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := json.Marshal(b.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{Body: data, Format: "JSON"}, nil
}

func (b *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}
