package Facade

import "testing"

func TestFacade(t *testing.T) {

	expect := "Build house\nTree grow\nChild born"

	man := newMan()

	result := man.MustDo()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
