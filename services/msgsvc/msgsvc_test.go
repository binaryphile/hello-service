package msgsvc

import "testing"

func TestWhenMessageThenReturnHelloWorld(t *testing.T) {
	expectedMessage := "Hello, world!"

	actualMessage := GetMessage()

	if actualMessage != expectedMessage {
		t.Errorf("GetMessage() == %q, want %q", actualMessage, expectedMessage)
	}
}
