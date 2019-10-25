package Bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {

	expect := "Suuuzzukkiii"

	car := NewCar(&EngineSuzuki{})

	sound := car.Rase()

	if sound != expect {
		t.Errorf("Expect sound to %s, but %s", expect, sound)
	}
}
