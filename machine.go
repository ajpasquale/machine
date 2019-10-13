package machine

import (
	"github.com/ajpasquale/timer"
)

// These constants represent different states of the machine
const (
	First = iota
	Second
)

// These constants represents different events of the machine
const (
	Transition = true
	Accepting  = false
)

type machine struct {
	state  int
	first  func()
	second func()
	timer  *timer.Timer
}

// NewMachine returns a binary state machine given two functions to execute
// during the accepting state and an interval to switch between.
func NewMachine(ff func(), sf func(), i int) *machine {
	m := &machine{
		state:  First,
		first:  ff,
		second: sf,
		timer:  timer.NewTimer(i),
	}
	return m
}

// Poll will return the machine state and decide to switch states or
// execute a function given by the user at creation.
func (m *machine) Poll() int {
	_, event := m.timer.Tick()
	switch m.state {
	case First:
		switch event {
		case Transition:
			m.state = Second
			break
		case Accepting:
			m.first()
			break
		}
		break
	case Second:
		switch event {
		case Transition:
			m.state = First
			break
		case Accepting:
			m.second()
			break
		}
		break
	}
	return m.state
}
