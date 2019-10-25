package Command

import (
	"testing"
)

func TestCommand(t *testing.T) {

	expect := "Toggle on!\n" +
		"Toggle off!\n"

	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})

	result := invoker.Execute()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s\n", expect, result)
	}
}
