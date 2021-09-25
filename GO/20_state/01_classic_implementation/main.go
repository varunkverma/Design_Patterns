package main

import "fmt"

// light: ON or OFF

type Switch struct {
	State State
}

func (sw *Switch) On() {
	sw.State.On(sw)
}

func (sw *Switch) Off() {
	sw.State.Off(sw)
}

func NewSwitch() *Switch {
	return &Switch{
		State: NewOffState(),
	}
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

// Provides no information, just the default behaviour for a specific implementation of State
type BaseState struct{}

func (cs *BaseState) On(s *Switch) {
	fmt.Println("Light is already ON")
}
func (cs *BaseState) Off(s *Switch) {
	fmt.Println("Light is already OFF")
}

// Technically, the only behvaiour that should be allowed on an ON state is to turn it OFF, but since, we need to make it as a valid type that implements the interface State, we rely on its embedded BaseState type to provide default functionality for turning it ON
type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned ON")
	return &OnState{
		BaseState: BaseState{},
	}
}

// So we make a unique Off() implementation of OnState
func (os *OnState) Off(sw *Switch) {
	fmt.Println("Turning the Light OFF")
	sw.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned OFF")
	return &OffState{
		BaseState: BaseState{},
	}
}

// So we make a unique On() implementation of OffState
func (os *OffState) On(sw *Switch) {
	fmt.Println("Turning the Light ON")
	sw.State = NewOnState()
}

func main() {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off()
}
