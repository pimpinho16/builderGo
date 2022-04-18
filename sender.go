package builderGo

type Sender struct {
	builder MessageBuilder
}

//SetBuilder asigna el constructor
func (s *Sender) SetBuilder(b MessageBuilder) {
	s.builder = b
}

//buildMessage
/*
Esto va a cambiar
*/
func (s *Sender) BuildMessage(recipient, message string) (*MessageRepresented, error) {
	return s.builder.SetRecipient(recipient).SetMessage(message).Build()

	/*
		s.builder.SetRecipient(recipient)
		s.builder.SetMessage(message)
		m, err := s.builder.Build()
		return m, err
	*/
	 
}
