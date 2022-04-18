package builderGo

import "testing"

func TestSender_BuildMessage(t *testing.T) {
	json := &JSONMessageBuilder{}
	sender := &Sender{}
	sender.SetBuilder(json)
	msg, err := sender.BuildMessage("Gopher", "Hola mundo")
	if err != nil {
		t.Fatalf("BuildMessage(): no se esperaba error, se obtuvo : %v", err)

	}
	t.Log(string(msg.Body))

	xml := &XMLJsonMessageBuilder{}

	sender.SetBuilder(xml)
	msg, err = sender.BuildMessage("Gopher", "Hola mundo")
	if err != nil {
		t.Fatalf("BuildMessage(): no se esperaba error, se obtuvo : %v", err)

	}
	t.Log(string(msg.Body))
}
