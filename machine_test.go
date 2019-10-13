package machine

import (
	"fmt"
	"testing"
)

func TestNewMachine(t *testing.T) {
	m := NewMachine(
		func() {
			fmt.Println("state one")
		},
		func() {
			fmt.Println("state two")
		},
		1,
	)
	if m == nil {
		t.Fatal("machine should not be nil")
	}
}

func TestPoll(t *testing.T) {
	m := NewMachine(
		func() {
			fmt.Println("state one")
		},
		func() {
			fmt.Println("state two")
		},
		1,
	)
	if m == nil {
		t.Fatal("machine should not be nil")
	}

	if m.Poll() != First {
		t.Errorf("machine is in the wrong state: %d", m.state)
	}

	for {
		if m.Poll() == Second {
			return
		}
	}
}
