package builderGo

//MessageBuilder
/*
esto utiliza el patron de cadena (no cadena de responsabilidad), por eso cada metodo devuelve la interfaz
*/
type MessageBuilder interface {
	SetRecipient(recipient string) MessageBuilder
	SetMessage(message string) MessageBuilder
	Build() (*MessageRepresented, error)
}
